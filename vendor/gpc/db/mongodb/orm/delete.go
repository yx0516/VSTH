package orm

import (
	"gopkg.in/mgo.v2"

	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm/structure"
)

//-----------------------------------------------------------------------------------------------------------//

// 删除
type Delete struct {
	Object interface{} // 结构体 对象
	IsAll  bool        // 是否删除 匹配的全部
	match  struct {    // 匹配 ，二选一，优先 M，未指定，默认会自动使用 id 字段
		M      interface{}
		Fields []string
	}
}

func NewDelete(object interface{}) *Delete {
	return &Delete{
		Object: object,
	}
}

// 设置 过滤 匹配 M
func (self *Delete) SetMatchM(m interface{}) *Delete {
	self.match.M = m
	return self
}

// 设置 过滤 匹配 字段
func (self *Delete) SetMatchFields(fields ...string) *Delete {
	self.match.Fields = append(self.match.Fields, fields...)
	return self
}

// 移除【只有 isAll 才会有 *mgo.ChangeInfo 值】
func (self *Delete) Remove() (*mgo.ChangeInfo, error) {
	si, err := structure.Ref(self.Object)
	if err != nil {
		return nil, err
	}

	if self.match.M == nil {
		if len(self.match.Fields) == 0 {
			self.match.Fields = []string{structure.CST_ID_NAME}
		}
		if self.match.M, err = si.FilterFields(true, self.match.Fields...); err != nil {
			return nil, err
		}
	}

	var info *mgo.ChangeInfo
	gsession.WithDB(func(db *mgo.Database) {
		coll := db.C(si.CollName)
		if self.IsAll {
			info, err = coll.RemoveAll(self.match.M)
		} else {
			err = coll.Remove(self.match.M)
		}
	})

	if err != nil && err == mgo.ErrNotFound {
		err = nil
	}
	return info, err
}

//-----------------------------------------------------------------------------------------------------------//

// 通过 ID 来删除
func DeleteId(objetc interface{}) error {
	_, err := NewDelete(objetc).SetMatchFields(structure.CST_ID_NAME).Remove()
	return err
}

//-----------------------------------------------------------------------------------------------------------//
