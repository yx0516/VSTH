package tag

import (
	"sync"

	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 缓存
type StructCache struct {
	cache   map[string]*StructTag
	convert ref.NameConvert
	tags    []string
	lock    sync.RWMutex
}

func NewStructCache(convert ref.NameConvert, tags ...string) *StructCache {
	return &StructCache{
		cache:   make(map[string]*StructTag),
		convert: convert,
		tags:    tags,
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 是否已经存在
func (self *StructCache) HasExist(structFulleName string) (ok bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	_, ok = self.cache[structFulleName]
	return
}

// 添加 缓存 或 获取缓存的值
func (self *StructCache) Add(object interface{}) (st *StructTag, err error) {
	var name string
	if name, err = ref.RefStructTypeFullName(object); err != nil {
		return
	}

	self.lock.Lock()
	defer self.lock.Unlock()
	cacheST, ok := self.cache[name]
	if !ok {
		cacheST = NewStructTag(self.convert, self.tags...)
		if err = cacheST.Reflect(object); err != nil {
			return
		}
		self.cache[name] = cacheST
	}

	st = NewStructTag(self.convert, self.tags...)
	st.Struct = &ref.Struct{
		FullName: cacheST.Struct.FullName,
		Name:     cacheST.Struct.Name,
		HasPtr:   cacheST.Struct.HasPtr,
		Fields:   make(map[string]*ref.StructField, len(cacheST.Struct.Fields)),
	}

	for name, sfi := range cacheST.Struct.Fields {
		st.Struct.Fields[name] = &ref.StructField{
			Name:  sfi.Name,
			Value: sfi.Value,
			Tags:  copyMap(sfi.Tags),
		}
	}
	st.Convert()

	if ok {
		err = st.Update(object)
	}

	return
}

func copyMap(src map[string]string) (dst map[string]string) {
	dst = make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
