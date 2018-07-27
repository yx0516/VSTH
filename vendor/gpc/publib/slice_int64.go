package publib

import (
	"errors"
	"fmt"
	"sort"
)

type Int64Array []int64

func (self Int64Array) Count() int { return len(self) }

func (self Int64Array) Sort()              { sort.Sort(self) }
func (self Int64Array) Len() int           { return len(self) }
func (self Int64Array) Less(i, j int) bool { return self[i] < self[j] }
func (self Int64Array) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func (self *Int64Array) Add(items ...int64) { *self = append(*self, items...) }

func (self Int64Array) Contains(item int64) (ok bool) {
	for _, v := range self {
		if v == item {
			return true
		}
	}
	return
}

func (self *Int64Array) Delete(item int64, isRecursion ...bool) (ok bool) {
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

func (self *Int64Array) Deletes(isRecursion bool, items ...int64) *Int64Array {
	for _, item := range items {
		self.Delete(item, isRecursion)
	}
	return self
}

func (self Int64Array) Subset(arr []int64) (err error) {
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

func (self *Int64Array) RemoveDuplicate() *Int64Array {
	found := make(map[int64]bool)
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

func (self Int64Array) Copy() Int64Array {
	cp := make(Int64Array, len(self))
	copy(cp, self)
	return cp
}

//-----------------------------------------------------------------------------------------------------------//
