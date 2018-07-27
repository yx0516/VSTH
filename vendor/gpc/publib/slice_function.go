package publib

import (
	"errors"
	"fmt"
	"strings"
)

// 数组里是否存在该子项
func StrArrayContains(arr *[]string, item string) (ok bool) {
	if arr != nil {
		for _, v := range *arr {
			if v == item {
				return true
			}
		}
	}
	return
}

// 数组里是否存在该子项
func IntArrayContains(arr *[]int, item int) (ok bool) {
	if arr != nil {
		for _, v := range *arr {
			if v == item {
				return true
			}
		}
	}
	return
}

// 数组里是否存在该子项
func UIntArrayContains(arr *[]uint, item uint) (ok bool) {
	if arr != nil {
		for _, v := range *arr {
			if v == item {
				return true
			}
		}
	}
	return
}

// 数组里是否存在该子项
func Int64ArrayContains(arr *[]int64, item int64) (ok bool) {
	if arr != nil {
		for _, v := range *arr {
			if v == item {
				return true
			}
		}
	}
	return
}

// 数组里是否存在该子项
func UInt64ArrayContains(arr *[]uint64, item uint64) (ok bool) {
	if arr != nil {
		for _, v := range *arr {
			if v == item {
				return true
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 数组里是否存在该子项 isRecursion=true 递归删除找到的所有，否则只删除找到的第一个【递归时返回的 ok 不正确】
func StrArrayDelete(arr *[]string, item string, isRecursion ...bool) (ok bool) {
	if arr != nil {
		count := len(*arr)
		loop := len(isRecursion) > 0 && isRecursion[0]
		for i := 0; i < count; i++ {
			if (*arr)[i] == item {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)
				ok = true
				if loop {
					i--
					count--
				} else {
					return
				}
			}
		}
	}
	return
}

// 删除 所有 元素 空字符串【包含 空空...】
func StrArrayDeleteAllSpace(arr *[]string) {
	if arr != nil {
		count := len(*arr)
		for i := 0; i < count; i++ {
			if strings.TrimSpace((*arr)[i]) == "" {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)

				i--
				count--
			}
		}
	}
}

// 所有 元素 TrimSpace
func StrArrayTrimSpace(arr *[]string) {
	count := len(*arr)
	for i := 0; i < count; i++ {
		(*arr)[i] = strings.TrimSpace((*arr)[i])
	}
}

// 删除 所有 元素 空字符串 并且 TrimSpace 【包含 空空...】
func StrArrayDeleteAllSpaceTrim(arr *[]string) {
	StrArrayDeleteAllSpace(arr)
	StrArrayTrimSpace(arr)
}

func StrArrayDeletes(arr *[]string, isRecursion bool, items ...string) {
	for _, item := range items {
		StrArrayDelete(arr, item, isRecursion)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 数组里是否存在该子项 isRecursion=true 递归删除找到的所有，否则只删除找到的第一个【递归时返回的 ok 不正确】
func IntArrayDelete(arr *[]int, item int, isRecursion ...bool) (ok bool) {
	if arr != nil {
		count := len(*arr)
		loop := len(isRecursion) > 0 && isRecursion[0]
		for i := 0; i < count; i++ {
			if (*arr)[i] == item {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)
				ok = true
				if loop {
					i--
					count--
				} else {
					return
				}
			}
		}
	}
	return
}

func IntArrayDeletes(arr *[]int, isRecursion bool, items ...int) {
	for _, item := range items {
		IntArrayDelete(arr, item, isRecursion)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 数组里是否存在该子项 isRecursion=true 递归删除找到的所有，否则只删除找到的第一个【递归时返回的 ok 不正确】
func UIntArrayDelete(arr *[]uint, item uint, isRecursion ...bool) (ok bool) {
	if arr != nil {
		count := len(*arr)
		loop := len(isRecursion) > 0 && isRecursion[0]
		for i := 0; i < count; i++ {
			if (*arr)[i] == item {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)
				ok = true
				if loop {
					i--
					count--
				} else {
					return
				}
			}
		}
	}
	return
}

func UIntArrayDeletes(arr *[]uint, isRecursion bool, items ...uint) {
	for _, item := range items {
		UIntArrayDelete(arr, item, isRecursion)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 数组里是否存在该子项 isRecursion=true 递归删除找到的所有，否则只删除找到的第一个【递归时返回的 ok 不正确】
func Int64ArrayDelete(arr *[]int64, item int64, isRecursion ...bool) (ok bool) {
	if arr != nil {
		count := len(*arr)
		loop := len(isRecursion) > 0 && isRecursion[0]
		for i := 0; i < count; i++ {
			if (*arr)[i] == item {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)
				ok = true
				if loop {
					i--
					count--
				} else {
					return
				}
			}
		}
	}
	return
}

func Int64ArrayDeletes(arr *[]int64, isRecursion bool, items ...int64) {
	for _, item := range items {
		Int64ArrayDelete(arr, item, isRecursion)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 数组里是否存在该子项 isRecursion=true 递归删除找到的所有，否则只删除找到的第一个【递归时返回的 ok 不正确】
func UInt64ArrayDelete(arr *[]uint64, item uint64, isRecursion ...bool) (ok bool) {
	if arr != nil {
		count := len(*arr)
		loop := len(isRecursion) > 0 && isRecursion[0]
		for i := 0; i < count; i++ {
			if (*arr)[i] == item {
				a1 := (*arr)[:i]
				a2 := (*arr)[i+1:]
				*arr = append(a1, a2...)
				ok = true
				if loop {
					i--
					count--
				} else {
					return
				}
			}
		}
	}
	return
}

func UInt64ArrayDeletes(arr *[]uint64, isRecursion bool, items ...uint64) {
	for _, item := range items {
		UInt64ArrayDelete(arr, item, isRecursion)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 检查 sArr 是否是 pArr 的子集
func StrArraySubset(pArr, sArr *[]string) (err error) {
	if pArr == nil || sArr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range *sArr {
		if !StrArrayContains(pArr, s) {
			return errors.New("field: " + s + " does not exist.")
		}
	}
	return
}

// 检查 sArr 是否是 pArr 的子集
func IntArraySubset(pArr, sArr *[]int) (err error) {
	if pArr == nil || sArr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range *sArr {
		if !IntArrayContains(pArr, s) {
			return fmt.Errorf("field: %v does not exist.", s)
		}
	}
	return
}

// 检查 sArr 是否是 pArr 的子集
func UIntArraySubset(pArr, sArr *[]uint) (err error) {
	if pArr == nil || sArr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range *sArr {
		if !UIntArrayContains(pArr, s) {
			return fmt.Errorf("field: %v does not exist.", s)
		}
	}
	return
}

// 检查 sArr 是否是 pArr 的子集
func Int64ArraySubset(pArr, sArr *[]int64) (err error) {
	if pArr == nil || sArr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range *sArr {
		if !Int64ArrayContains(pArr, s) {
			return fmt.Errorf("field: %v does not exist.", s)
		}
	}
	return
}

// 检查 sArr 是否是 pArr 的子集
func UInt64ArraySubset(pArr, sArr *[]uint64) (err error) {
	if pArr == nil || sArr == nil {
		return errors.New("args is nil.")
	}
	for _, s := range *sArr {
		if !UInt64ArrayContains(pArr, s) {
			return fmt.Errorf("field: %v does not exist.", s)
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 移除重复的 string slice 值
func StrArrayRemoveDuplicate(sliceString *[]string) {
	found := make(map[string]bool)
	total := 0
	for i, val := range *sliceString {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*sliceString)[total] = (*sliceString)[i]
			total++
		}
	}
	*sliceString = (*sliceString)[:total]
}

// 移除重复的 int slice 值
func IntArrayRemoveDuplicate(sliceInt *[]int) {
	found := make(map[int]bool)
	total := 0
	for i, val := range *sliceInt {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*sliceInt)[total] = (*sliceInt)[i]
			total++
		}
	}
	*sliceInt = (*sliceInt)[:total]
}

// 移除重复的 uint slice 值
func UIntArrayRemoveDuplicate(sliceUInt *[]uint) {
	found := make(map[uint]bool)
	total := 0
	for i, val := range *sliceUInt {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*sliceUInt)[total] = (*sliceUInt)[i]
			total++
		}
	}
	*sliceUInt = (*sliceUInt)[:total]
}

// 移除重复的 int64 slice 值
func Int64ArrayRemoveDuplicate(sliceInt64 *[]int64) {
	found := make(map[int64]bool)
	total := 0
	for i, val := range *sliceInt64 {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*sliceInt64)[total] = (*sliceInt64)[i]
			total++
		}
	}
	*sliceInt64 = (*sliceInt64)[:total]
}

// 移除重复的 UInt64 slice 值
func UInt64ArrayRemoveDuplicate(sliceUInt64 *[]uint64) {
	found := make(map[uint64]bool)
	total := 0
	for i, val := range *sliceUInt64 {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*sliceUInt64)[total] = (*sliceUInt64)[i]
			total++
		}
	}
	*sliceUInt64 = (*sliceUInt64)[:total]
}

//-----------------------------------------------------------------------------------------------------------//
