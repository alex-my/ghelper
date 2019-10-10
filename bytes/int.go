package bytes

import (
	"encoding/binary"
)

// ToInt8 转为 int8
func ToInt8(p []byte) int8 {
	return int8(p[0])
}

// PutInt8 int8 转 []byte
func PutInt8(n int8, p []byte) {
	_ = p[0]
	p[0] = byte(n)
}

// ToUint8 转为 uint8
func ToUint8(p []byte) uint8 {
	return uint8(p[0])
}

// PutUint8 uint8 转 []byte
func PutUint8(n uint8, p []byte) {
	_ = p[0]
	p[0] = byte(n)
}

// ToInt16 转 int16
func ToInt16(p []byte) int16 {
	return int16(binary.BigEndian.Uint16(p))
}

// PutInt16 int16 转 []byte
func PutInt16(n int16, p []byte) {
	binary.BigEndian.PutUint16(p, uint16(n))
}

// ToUint16 转 uint16
func ToUint16(p []byte) uint16 {
	return binary.BigEndian.Uint16(p)
}

// PutUint16 uint16 转 []byte
func PutUint16(n uint16, p []byte) {
	binary.BigEndian.PutUint16(p, n)
}

// ToInt32 转 int32
func ToInt32(p []byte) int32 {
	return int32(binary.BigEndian.Uint32(p))
}

// PutInt32 int32 转 []byte
func PutInt32(n int32, p []byte) {
	binary.BigEndian.PutUint32(p, uint32(n))
}

// ToUint32 转 uint32
func ToUint32(p []byte) uint32 {
	return binary.BigEndian.Uint32(p)
}

// PutUint32 uint32 转 []byte
func PutUint32(n uint32, p []byte) {
	binary.BigEndian.PutUint32(p, n)
}

// ToInt64 转 int64
func ToInt64(p []byte) int64 {
	return int64(binary.BigEndian.Uint64(p))
}

// PutInt64 int64 转 []byte
func PutInt64(n int64, p []byte) {
	binary.BigEndian.PutUint64(p, uint64(n))
}

// ToUint64 转 uint64
func ToUint64(p []byte) uint64 {
	return binary.BigEndian.Uint64(p)
}

// PutUint64 uint64 转 []byte
func PutUint64(n uint64, p []byte) {
	binary.BigEndian.PutUint64(p, n)
}
