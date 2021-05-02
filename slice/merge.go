package slice

import (
	"strconv"
	"strings"
)

// SliceMerge merges interface slices to one slice.
func Merge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

func MergeInt(slice1, slice2 []int) (c []int) {
	c = append(slice1, slice2...)
	return
}

func MergeInt64(slice1, slice2 []int64) (c []int64) {
	c = append(slice1, slice2...)
	return
}

func MergeString(slice1, slice2 []string) (c []string) {
	c = append(slice1, slice2...)
	return
}

// 合并数组
func MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

// 删除数组
func DeleteArrayElm(src []interface{}, index int) (result []interface{}) {
	result = append(src[:index], src[(index+1):]...)
	return
}

// []string => []int
func ArrayStr2Int(data []string) []int {
	var (
		arr = make([]int, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i := range data {
		var num, _ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func ArrayInt2Str(data []int) []string {

	arr := make([]string, 0, len(data))

	if len(data) == 0 {
		return arr
	}
	for i := range data {
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

// str[TrimSpace] in string list
func TrimSpaceStrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == strings.TrimSpace(row) {
				return true
			}
		}
	}
	return false
}

// str in string list
func StrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == row {
				return true
			}
		}
	}
	return false
}

// num in int list
func IntInArray(num int, data []int) bool {
	if len(data) > 0 {
		for _, row := range data {
			if num == row {
				return true
			}
		}
	}
	return false
}
