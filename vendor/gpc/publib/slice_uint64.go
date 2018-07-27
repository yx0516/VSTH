package publib

import (
	"errors"
	"fmt"
	"sort"
)

type UInt64Array []uint64

func (self UInt64Array) Count() int { return len(self) }

func (self UInt64Array) Sort()              { sort.Sort(self) }
func (self UInt64Array) Len() int           { return len(self) }
func (self UInt64Array) Less(i, j int) bool { return self[i] < self[j] }
func (self UInt64Array) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *UInt64Array) Add(items ...uint64) { *self = append(*self, items...) }

func (self UInt64Array) Contains(item uint64) (ok bool) {
	for _, v := range self {
		if v == item {
			return true
		}
	}
	return
}

func (self *UInt64Array) Delete(item uint64, isRecursion ...bool) (ok bool) {
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

func (self *UInt64Array) Deletes(isRecursion bool, items ...uint64) *UInt64Array {
	for _, item := range items {
		self.Delete(item, isRecursion)
	}
	return self
}

func (self UInt64Array) Subset(arr []uint64) (err error) {
	if arr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range arr {
		if !self.Contains(s) {
			return fmt.Errorf("field %v does not exist.", s)
		}
	}
	return
}

func (self *UInt64Array) RemoveDuplicate() *UInt64Array {
	found := make(map[uint64]bool)
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

func (self UInt64Array) Copy() UInt64Array {
	cp := make(UInt64Array, len(self))
	copy(cp, self)
	return cp
}

//-----------------------------------------------------------------------------------------------------------//
