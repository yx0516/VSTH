package orm

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm/structure"
)

//-----------------------------------------------------------------------------------------------------------//

// 更新
type Update struct {
	Object    interface{} // 结构体 对象
	IsAll     bool        // 是否更新 匹配的全部
	IsUpsert  bool        // 是否 不存在，则插入
	IsReplace bool        // 是否 替换，而不是 $set
	Match     interface{} // 匹配
	Set       struct {    // 更新 ，二选一，优先 M
		M      bson.M
		Fields []string
	}
}

func NewUpdate(object interface{}) *Update {
	return &Update{
		Object: object,
	}
}

// 设置 更新内容 M
func (self *Update) SetM(m bson.M) *Update {
	self.Set.M = m
	return self
}

// 设置 更新 字段
func (self *Update) SetFields(fields ...string) *Update {
	self.Set.Fields = append(self.Set.Fields, fields...)
	return self
}

// 更新
func (self *Update) Update() (*mgo.ChangeInfo, error) {
	if self.Set.M == nil && len(self.Set.Fields) == 0 {
		return nil, nil
	}

	si, err := structure.Ref(self.Object)
	if err != nil {
		return nil, err
	}

	if self.Match == nil {
		if m, err := si.GetId(); err != nil {
			return nil, err
		} else {
			self.Match = m
		}
	}

	if self.Set.M == nil {
		if self.Set.M, err = si.FilterFields(true, self.Set.Fields...); err != nil {
			return nil, err
		}
	}
	delete(self.Set.M, structure.CST_ID_NAME)

	var info *mgo.ChangeInfo
	gsession.WithDB(func(db *mgo.Database) {
		coll := db.C(si.CollName)

		data := bson.M{"$set": self.Set.M}
		if !self.IsAll && self.IsReplace {
			data = self.Set.M
		}

		if self.IsUpsert {
			info, err = coll.Upsert(self.Match, data)
		} else if self.IsAll {
			info, err = coll.UpdateAll(self.Match, data)
		} else {
			err = coll.Update(self.Match, data)
		}
	})

	return info, err
}

//-----------------------------------------------------------------------------------------------------------//
