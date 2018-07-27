package publib

import (
	"errors"
	"sort"
	"strings"
)

type StringArray []string

func (self *StringArray) Sort() *StringArray {
	sort.Sort(self)
	return self
}

func (self StringArray) Len() int           { return len(self) }
func (self StringArray) Less(i, j int) bool { return self[i] < self[j] }
func (self *StringArray) Swap(i, j int) {
	(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
}

func (self StringArray) Count() int { return len(self) }

func (self *StringArray) Add(items ...string) *StringArray {
	*self = append(*self, items...)
	return self
}

func (self StringArray) Contains(item string) (ok bool) {
	for _, v := range self {
		if v == item {
			return true
		}
	}
	return
}

func (self StringArray) ContainsSub(item string) (ok bool) {
	for _, v := range self {
		if strings.Contains(item, v) {
			return true
		}
	}
	return
}

func (self *StringArray) Delete(item string, isRecursion ...bool) (ok bool) {
	count := self.Len()
	loop := len(isRecursion) > 0 && isRecursion[0]
	for i := 0; i < count; i++ {
		if (*self)[i] == item {
			a1 := (*self)[:i]
			a2 := (*self)[i+1:]
			*self = append(a1, a2...)
			ok = true
			if loop {
				i--
				count--
			} else {
				return
			}
		}
	}
	return
}

func (self *StringArray) DeleteSpace() *StringArray {
	if count := self.Count(); count > 0 {
		var arr []string
		for i := 0; i < count; i++ {
			if (*self)[i] != "" {
				arr = append(arr, (*self)[i])
			}
		}
		*self = arr
	}
	return self
}

func (self *StringArray) Trim(cutset string) *StringArray {
	count := self.Count()
	for i := 0; i < count; i++ {
		(*self)[i] = strings.Trim((*self)[i], cutset)
	}
	return self
}

func (self *StringArray) TrimDeleteSpace(cutset string) *StringArray {
	if count := self.Count(); count > 0 {
		var arr []string
		for i := 0; i < count; i++ {
			if (*self)[i] = strings.Trim((*self)[i], cutset); (*self)[i] != "" {
				arr = append(arr, (*self)[i])
			}
		}
		*self = arr
	}
	return self
}

func (self *StringArray) TrimSpace() *StringArray {
	count := self.Count()
	for i := 0; i < count; i++ {
		(*self)[i] = strings.TrimSpace((*self)[i])
	}
	return self
}

func (self *StringArray) TrimSpaceDeleteSpace() *StringArray {
	if count := self.Count(); count > 0 {
		var arr []string
		for i := 0; i < count; i++ {
			if (*self)[i] = strings.TrimSpace((*self)[i]); (*self)[i] != "" {
				arr = append(arr, (*self)[i])
			}
		}
		*self = arr
	}
	return self
}

func (self *StringArray) Deletes(isRecursion bool, items ...string) *StringArray {
	for _, item := range items {
		self.Delete(item, isRecursion)
	}
	return self
}

func (self StringArray) Subset(arr []string) (err error) {
	if arr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range arr {
		if !self.Contains(s) {
			return errors.New("field: " + s + " does not exist.")
		}
	}
	return
}

func (self *StringArray) RemoveDuplicate() *StringArray {
	found := make(map[string]bool)
	total := 0
	for i, val := range *self {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*self)[total] = (*self)[i]
			total++
		}
	}
	*self = (*self)[:total]
	return self
}

func (self StringArray) Copy() StringArray {
	cp := make(StringArray, len(self))
	copy(cp, self)
	return cp
}

func (self StringArray) Join(sep string) string {
	return strings.Join(self, sep)
}

func (self StringArray) ToString() string {
	return strings.Join(self, "|")
}

func (self StringArray) ToLower() {
	count := self.Len()
	for i := 0; i < count; i++ {
		self[i] = strings.ToLower(self[i])
	}
}

func (self StringArray) ToUpper() {
	count := self.Len()
	for i := 0; i < count; i++ {
		self[i] = strings.ToUpper(self[i])
	}
}

//-----------------------------------------------------------------------------------------------------------//
