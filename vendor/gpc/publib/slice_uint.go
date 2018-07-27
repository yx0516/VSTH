package publib

import (
	"errors"
	"fmt"
	"sort"
)

type UIntArray []uint

func (self UIntArray) Count() int { return len(self) }

func (self UIntArray) Sort()              { sort.Sort(self) }
func (self UIntArray) Len() int           { return len(self) }
func (self UIntArray) Less(i, j int) bool { return self[i] < self[j] }
func (self UIntArray) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *UIntArray) Add(items ...uint) { *self = append(*self, items...) }

func (self UIntArray) Contains(item uint) (ok bool) {
	for _, v := range self {
		if v == item {
			return true
		}
	}
	return
}

func (self *UIntArray) Delete(item uint, isRecursion ...bool) (ok bool) {
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

func (self *UIntArray) Deletes(isRecursion bool, items ...uint) *UIntArray {
	for _, item := range items {
		self.Delete(item, isRecursion)
	}
	return self
}

func (self UIntArray) Subset(arr []uint) (err error) {
	if arr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range arr {
		if !self.Contains(s) {
			return fmt.Errorf("field: %v does not exist.", s)
		}
	}
	return
}

func (self *UIntArray) RemoveDuplicate() *UIntArray {
	found := make(map[uint]bool)
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

func (self UIntArray) Copy() UIntArray {
	cp := make(UIntArray, len(self))
	copy(cp, self)
	return cp
}

//-----------------------------------------------------------------------------------------------------------//
