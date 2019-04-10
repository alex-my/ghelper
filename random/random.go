package random

import (
	_cr "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String 获取指定长度的字符串
func String(length int) string {
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

// Int 获取指定范围内的整数
func Int(min, max int64) int64 {
	if min >= max || min == max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// Uint32 获取随机数 uint32
func Uint32() uint32 {
	var v uint32
	if err := binary.Read(_cr.Reader, binary.BigEndian, &v); err == nil {
		return v
	}
	panic("Random failed")
}
