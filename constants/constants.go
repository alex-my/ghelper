// Package constants 定义一些常量
package constants

const (
	// Uint16Max 65535
	Uint16Max = ^uint16(0)

	// Uint16Min 0
	Uint16Min uint16 = 0

	// Uint32Max 4294967295
	Uint32Max = ^uint32(0)

	// Uint32Min 0
	Uint32Min uint32 = 0

	// Uint64Max 18446744073709551615
	Uint64Max = ^uint64(0)

	// Uint64Min 0
	Uint64Min = 0

	// Int16Max 32767
	Int16Max = int16(^uint16(0) >> 1)

	// Int16Min -32768
	Int16Min = ^Int16Max

	// Int32Max 2147483647
	Int32Max = int32(^uint32(0) >> 1)

	// Int32Min -2147483648
	Int32Min = ^Int32Max

	// Int64Max 9223372036854775807
	Int64Max = int64(^uint64(0) >> 1)

	// Int64Min -9223372036854775808
	Int64Min = ^Int64Max
)
