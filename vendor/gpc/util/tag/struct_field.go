package tag

import (
	"errors"
	"strings"
	"sync"

	"gpc/util/ref"

	plib "gpc/publib"

	deep "gpc/util/deep/json"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 属性
type StructField struct {
	Xorm    MapSSA
	White   plib.StringArray
	Black   plib.StringArray
	Query   plib.StringArray
	All     plib.StringArray
	Fileds  plib.StringArray
	ZeroNil plib.StringArray
	NilDel  plib.StringArray
}

// 如果 filterQuery == true ，则只获取指定 xorm:="query" 属性的才是查询字段， 否则默认全部字段可查询
func NewStructField(st *StructTag, filterQuery ...bool) *StructField {
	if st != nil {
		sf := &StructField{
			Xorm:   make(MapSSA),
			All:    st.GetFieldNames(),
			Fileds: st.GetStrucFieldNames(),
		}

		for field, info := range st.Fields {
			if tag, ok := info.Tags[CST_TAG_XORM]; ok {
				vals := plib.StringArray(strings.Split(tag, ";"))
				for _, val := range vals {
					val = strings.TrimSpace(val)
					switch val {
					case CST_XORM_ZERO_NIL:
						sf.ZeroNil.Add(field)
						if vals.Contains(CST_XORM_ADD_FK) {
							sf.ZeroNil.Add(info.Name)
						}
					case CST_XORM_NIL_DEL:
						sf.NilDel.Add(field)
						if vals.Contains(CST_XORM_ADD_FK) {
							sf.NilDel.Add(info.Name)
						}
					}
					sf.Xorm[field] = append(sf.Xorm[field], val)
				}
			}
		}

		for field, arr := range sf.Xorm {
			if arr.Contains(CST_XORM_UPDATE) {
				sf.White.Add(field)
			}
		}

		for _, field := range sf.All {
			if !sf.White.Contains(field) {
				sf.Black.Add(field)
			}
		}

		if len(filterQuery) > 0 && filterQuery[0] {
			for field, arr := range sf.Xorm {
				if arr.Contains(CST_XORM_QUERY) {
					sf.Query.Add(field)
				}
			}
		} else {
			sf.Query = st.GetFieldNames()
		}

		return sf
	} else {
		return &StructField{}
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 管理 结构体 属性
type AdminStructField struct {
	sfs map[string]*StructField
	mt  sync.RWMutex
}

func NewAdminStructField() *AdminStructField {
	return &AdminStructField{
		sfs: make(map[string]*StructField),
	}
}

func (self *AdminStructField) HasExist(structFulleName string) (ok bool) {
	self.mt.Lock()
	defer self.mt.Unlock()

	_, ok = self.sfs[structFulleName]
	return
}

func (self *AdminStructField) Add(structFulleName string, sf *StructField) {
	self.mt.Lock()
	defer self.mt.Unlock()

	self.sfs[structFulleName] = sf
}

func (self *AdminStructField) Get(v interface{}) (*StructField, error) {
	if v == nil {
		return nil, errors.New("key is nil.")
	}

	self.mt.RLock()
	defer self.mt.RUnlock()

	if key, ok := v.(string); ok {
		return self.sfs[key], nil
	}

	if key, err := ref.RefStructTypeFullName(v); err != nil {
		return nil, err
	} else if sf, ok := self.sfs[key]; ok {
		return sf, nil
	} else {
		return nil, errors.New("key <" + key + "> not fount.")
	}
}

func (self *AdminStructField) GetCopy(v interface{}) (*StructField, error) {
	if sf, err := self.Get(v); err == nil {
		cp := new(StructField)
		err = deep.Copy(sf, cp)
		return cp, err
	} else {
		return nil, err
	}
}

//-----------------------------------------------------------------------------------------------------------//
