package array

import (
	"strconv"
	"strings"
)

// IntToString int 数组转字符串
// [1,2,3,4] -> "1,2,3,4"
func IntToString(i []int) string {
	l := make([]string, 0, len(i))
	for _, v := range i {
		l = append(l, strconv.Itoa(v))
	}
	return strings.Join(l, ",")
}

// Int64ToString int64 数组转字符串
// [1,2,3,4] -> "1,2,3,4"
func Int64ToString(i64 []int64) string {
	l := make([]string, 0, len(i64))
	for _, v := range i64 {
		l = append(l, strconv.FormatInt(v, 10))
	}
	return strings.Join(l, ",")
}

// StringToString 字符数组转字符串
// ["1","2","3","4"] -> "1,2,3,4"
func StringToString(s []string) string {
	return strings.Join(s, ",")
}
