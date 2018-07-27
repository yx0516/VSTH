package publib

import (
	"errors"
	"fmt"
	"sort"
)

type IntArray []int

func (self IntArray) Count() int { return len(self) }

func (self IntArray) Sort()              { sort.Sort(self) }
func (self IntArray) Len() int           { return len(self) }
func (self IntArray) Less(i, j int) bool { return self[i] < self[j] }
func (self IntArray) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *IntArray) Add(items ...int) { *self = append(*self, items...) }

func (self IntArray) Contains(item int) (ok bool) {
	for _, v := range self {
		if v == item {
			return true
		}
	}
	return
}

func (self *IntArray) Delete(item int, isRecursion ...bool) (ok bool) {
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

func (self *IntArray) Deletes(isRecursion bool, items ...int) *IntArray {
	for _, item := range items {
		self.Delete(item, isRecursion)
	}
	return self
}

func (self IntArray) Subset(arr []int) (err error) {
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

func (self *IntArray) RemoveDuplicate() *IntArray {
	found := make(map[int]bool)
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

func (self IntArray) Copy() IntArray {
	cp := make(IntArray, len(self))
	copy(cp, self)
	return cp
}

//-----------------------------------------------------------------------------------------------------------//
