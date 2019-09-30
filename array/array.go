package array

import (
	"strconv"
	"strings"
)

// IntToString int 数组转字符串
// eg: IntToString([]int{1, 2, 3, 4}) 		-> "1,2,3,4"
// eg: IntToString([]int{1, 2, 3, 4}, "+") 	-> "1+2+3+4"
func IntToString(i []int, sep ...string) string {
	l := make([]string, 0, len(i))
	for _, v := range i {
		l = append(l, strconv.Itoa(v))
	}
	s := ","
	if len(sep) > 0 && sep[0] != "" {
		s = sep[0]
	}
	return strings.Join(l, s)
}

// Int64ToString int64 数组转字符串
// eg: Int64ToString([]int64{1, 2, 3, 4})		-> "1,2,3,4"
// eg: Int64ToString([]int64{1, 2, 3, 4}, "+")	-> "1+2+3+4"
func Int64ToString(i64 []int64, sep ...string) string {
	l := make([]string, 0, len(i64))
	for _, v := range i64 {
		l = append(l, strconv.FormatInt(v, 10))
	}
	s := ","
	if len(sep) > 0 && sep[0] != "" {
		s = sep[0]
	}
	return strings.Join(l, s)
}

// StringToString 字符数组转字符串
// eg: StringToString(["1","2","3","4"]) 		-> "1,2,3,4"
// eg: StringToString(["1","2","3","4"], "+") 	-> "1+2+3+4"
func StringToString(str []string, sep ...string) string {
	s := ","
	if len(sep) > 0 && sep[0] != "" {
		s = sep[0]
	}
	return strings.Join(str, s)
}
