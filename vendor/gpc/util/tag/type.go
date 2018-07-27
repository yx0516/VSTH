package tag

import (
	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// Map 字符串数组
type MapSSA map[string]plib.StringArray

func (self MapSSA) IncludeItems(items ...string) (keys plib.StringArray) {
	if len(items) == 0 {
		return
	}
	var ok bool
	for key, arr := range self {
		ok = true
		for _, item := range items {
			if ok = arr.Contains(item); !ok {
				break
			}
		}
		if ok {
			keys.Add(key)
		}
	}
	return
}

func (self MapSSA) ExcludeItems(items ...string) (keys plib.StringArray) {
	if len(items) == 0 {
		return
	}
	var ok bool
	for key, arr := range self {
		ok = false
		for _, item := range items {
			if ok = arr.Contains(item); ok {
				break
			}
		}
		if !ok {
			keys.Add(key)
		}
	}
	return
}

func (self MapSSA) Contains(key string) (ok bool) {
	_, ok = self[key]
	return
}

func (self MapSSA) Len() int {
	return len(self)
}

func (self MapSSA) Delete(keys ...string) {
	for _, key := range keys {
		delete(self, key)
	}
}

func (self MapSSA) Retain(keys ...string) {
	delKeys := self.Keys()
	delKeys.Deletes(true, keys...)
	self.Delete(delKeys...)
}

func (self MapSSA) Keys() (keys plib.StringArray) {
	for key := range self {
		keys.Add(key)
	}
	return
}

func (self MapSSA) GetMap(keys ...string) (m MapSSA) {
	m = make(MapSSA)
	for _, key := range keys {
		if val, ok := self[key]; ok {
			m[key] = val
		}
	}
	return
}

func (self MapSSA) Add(key string, val plib.StringArray) {
	self[key] = val
}

func (self MapSSA) Append(m map[string][]string) {
	for key, val := range m {
		self[key] = val
	}
}

func (self MapSSA) Clear() {
	self.Delete(self.Keys()...)
}

//-----------------------------------------------------------------------------------------------------------//
