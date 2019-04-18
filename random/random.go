package random

import (
	"bytes"
	_cr "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var bufferPool *sync.Pool

func init() {
	rand.Seed(time.Now().UnixNano())

	bufferPool = &sync.Pool{}
	bufferPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

// buffer 从池中获取 buffer
func buffer() *bytes.Buffer {
	buff := bufferPool.Get().(*bytes.Buffer)
	buff.Reset()
	return buff
}

// releaseBuffer 将 buff 放入池中
func releaseBuffer(buff *bytes.Buffer) {
	bufferPool.Put(buff)
}

// String 获取指定长度的字符串
func String(length int) string {
	buf := buffer()
	defer releaseBuffer(buf)

	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			buf.WriteString(strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			buf.WriteString(string(rand.Intn(26) + 65))
		} else {
			buf.WriteString(string(rand.Intn(26) + 97))
		}
	}
	return buf.String()
}

// Int 获取指定范围内的整数
func Int(min, max int64) int64 {
	if min >= max || min == max {
		return max
	}
	return rand.Int63n(max-min) + min
}

// Uint32 获取随机数，类型为 uint32
func Uint32() uint32 {
	var v uint32
	if err := binary.Read(_cr.Reader, binary.BigEndian, &v); err == nil {
		return v
	}
	panic("Random failed")
}
