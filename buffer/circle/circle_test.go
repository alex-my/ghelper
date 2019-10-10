package circle_test

import (
	"testing"

	"github.com/alex-my/ghelper/buffer/circle"
)

func TestCircleWrite(t *testing.T) {
	c := circle.New(10)
	n, err := c.Write([]byte("01234"))
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 5 {
		t.Fatalf("n must be 5, now: %d", n)
	}

	_, err = c.Write([]byte("56789"))
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}

	if c.Len() != 10 {
		t.Fatalf("c.len must be 10, now: %d", c.Len())
	}

	if c.Free() != 0 {
		t.Fatalf("c.free must be 0, now: %d", c.Free())
	}
}

func TestCircleRead(t *testing.T) {
	c := circle.New(10)
	c.Write([]byte("01234"))

	// 只读取部分
	buffer := make([]byte, 3)
	n, err := c.Read(buffer)
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 3 {
		t.Fatalf("n must be 3, now: %d", n)
	}
	if c.Len() != 2 {
		t.Fatalf("c.len must be 2, now: %d", c.Len())
	}
	if c.Free() != 8 {
		t.Fatalf("c.free must be 8, now: %d", c.Free())
	}

	// 将剩余部分全部读取
	buffer = make([]byte, 10)
	n, err = c.Read(buffer)
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 2 {
		t.Fatalf("n must be 2, now: %d", n)
	}
}

func TestCircleFull(t *testing.T) {
	c := circle.New(10)

	// 写入 9 个数据
	c.Write([]byte("012345678"))
	t.Log(c)

	// 先读取 8 个数据，还剩下 1 个数据
	buffer := make([]byte, 8)
	n, err := c.Read(buffer)
	t.Log(c, buffer)

	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 8 {
		t.Fatalf("n must be 8, now: %d", n)
	}

	// 写入 5 个数据
	n, err = c.Write([]byte("abcde"))
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 5 {
		t.Fatalf("n must be 5, now: %d", n)
	}

	// 读取 5 个数据
	buffer = make([]byte, 5)
	n, err = c.Read(buffer)
	t.Log(c, buffer)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 5 {
		t.Fatalf("n must be 5, now: %d", n)
	}

	// 写入 8 个数据
	n, err = c.Write([]byte("fghijklm"))
	t.Log(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n != 8 {
		t.Fatalf("n must be 8, now: %d", n)
	}
}

func TestCircleReadN(t *testing.T) {
	c := circle.New(10)
	c.Write([]byte("012345"))

	buffer := make([]byte, 2)
	err := c.ReadN(3, buffer)
	if err != circle.ErrInvalidBuffer {
		t.Fatal(err.Error())
	}

	buffer = make([]byte, 5)
	err = c.ReadN(3, buffer)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(c, buffer)
}

// go tool pprof circle.test profile_cpu.out

// go test -test.bench="BenchmarkWriteAndRead1" -benchtime=5s -cpuprofile profile_cpu.out
// BenchmarkWriteAndRead1-12
// 446801310	        13.2 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/alex-my/ghelper/buffer/circle	7.452s
func BenchmarkWriteAndRead1(b *testing.B) {
	b.ReportAllocs()

	c := circle.New(10)
	buffer := make([]byte, 5)

	for i := 0; i < b.N; i++ {
		c.Write([]byte("12345"))
		c.Read(buffer)
	}
}

// go test -test.bench="BenchmarkWriteAndRead2" -benchtime=5s -cpuprofile profile_cpu.out
// BenchmarkWriteAndRead2-12
// 208808558	        28.6 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/alex-my/ghelper/buffer/circle	9.072s
func BenchmarkWriteAndRead2(b *testing.B) {
	b.ReportAllocs()

	c := circle.New(10)
	buffer1 := make([]byte, 5)
	buffer2 := make([]byte, 7)

	for i := 0; i < b.N; i++ {
		c.Write([]byte("1234567"))
		c.Read(buffer1)

		c.Write([]byte("89012"))
		c.Read(buffer2)
	}
}
