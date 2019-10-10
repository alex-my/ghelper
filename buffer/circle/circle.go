package circle

import (
	"errors"
	"fmt"
)

// 环形缓冲区
// 实现了 io.Reader, io.Writer, io.ReadWriter 接口
// 非线程安全

var (
	// ErrTooManyToWrite 内容太长，无法写入
	ErrTooManyToWrite = errors.New("too many data to write")
	// ErrInvalidBuffer 用于读取的 buffer 无效
	ErrInvalidBuffer = errors.New("invalid buffer")
	// ErrNotEnough 数据不足以读取长度 n 的内容
	ErrNotEnough = errors.New("not enough data")
	// ErrInvalidSkipSize 无效的长度
	ErrInvalidSkipSize = errors.New("invalid skip size")
	// ErrIsFull 缓冲区已满
	ErrIsFull = errors.New("buffer is full")
	// ErrIsEmpty 缓冲区已满
	ErrIsEmpty = errors.New("buffer is empty")
)

// Circle 环形结构
type Circle struct {
	write int
	read  int
	size  int

	buffer []byte
	isFull bool
}

// New 创建环形结构
func New(size int) *Circle {
	if size <= 0 {
		panic("size must be bigger than zero")
	}

	return &Circle{
		write:  0,
		read:   0,
		size:   size,
		buffer: make([]byte, size),
		isFull: false,
	}
}

// Reset 重置
func (c *Circle) Reset() {
	c.write = 0
	c.read = 0
	c.isFull = false
}

// Len 有效 buffer 的长度
func (c *Circle) Len() int {
	// read 到 write 中间存储着有效的 buffer

	if c.isFull {
		return c.size
	}

	// |___read______write____|
	if c.write >= c.read {
		return c.write - c.read
	}

	// |___write______read____|
	return c.size - (c.read - c.write)
}

// Cap ..
func (c *Circle) Cap() int {
	return c.size
}

// Free 剩余空间
func (c *Circle) Free() int {
	return c.size - c.Len()
}

// Write 写入数据
// 只有全部写入，或者都不写入两种情况，不存在只写入一部分的情形
func (c *Circle) Write(p []byte) (int, error) {
	if p == nil || len(p) == 0 {
		return 0, nil
	}

	if c.isFull {
		return 0, ErrIsFull
	}

	free := c.Free()
	lenp := len(p)
	if lenp > free {
		return 0, ErrTooManyToWrite
	}

	// 空闲部分足以写入数据
	// 开始写入数据

	// start|___read______write____|end
	if c.write >= c.read {
		writeToEnd := c.size - c.write

		// 如果 write 到 end 能写入全部数据
		if writeToEnd >= lenp {
			copy(c.buffer[c.write:], p)
			c.write += lenp
		} else {
			// 数据分两部分写入
			// 第一部分写入 write 到 end 的空间
			copy(c.buffer[c.write:], p[:writeToEnd])
			// 第二部分写到 start 到 read 之间
			copy(c.buffer[:], p[writeToEnd:])
			left := lenp - writeToEnd

			// start|__write_read__________|end
			c.write = left
		}
	} else {
		// start|___write______read____|end
		copy(c.buffer[c.write:], p)
		c.write += lenp
	}

	if c.write == c.read {
		c.isFull = true
	}

	return lenp, nil
}

// Read 读取尽量多的数据到 p 中
// p: 数据会拷贝到 p 中
// n: 读取到 P 中的数据长度
func (c *Circle) Read(p []byte) (int, error) {
	if p == nil || len(p) == 0 {
		return 0, ErrInvalidBuffer
	}

	// 没有内容
	if c.IsEmpty() {
		return 0, ErrIsEmpty
	}

	lenp := len(p)

	// start|___read______write____|end
	// 可以读取 read 到 write 之间的数据
	if c.write > c.read {
		// 获取有效 buffer 长度
		readToWrite := c.write - c.read

		if readToWrite > lenp {
			// p 存储不下，只能拷贝部分
			readToWrite = lenp
		}

		copy(p, c.buffer[c.read:c.read+readToWrite])

		// read 向前移动
		c.read += readToWrite

		c.isFull = false
		return readToWrite, nil
	}

	// c.write <= c.read
	// start|___write______read____|end
	// 可以读取两部分数据
	// 第一部分 read 到 end
	// 第二部分 start 到 write

	canReadLen := c.Len()
	if canReadLen > lenp {
		// 内容太长，p 无法全部读取
		canReadLen = lenp
	}

	if c.read+canReadLen <= c.size {
		// read 到 end 的内容就满足 p 的读取需求
		copy(p, c.buffer[c.read:c.read+canReadLen])

		// start|___write________read__|end
		c.read += canReadLen
	} else {
		// 两部分都需要发生拷贝行为

		// 拷贝 read 到 end 的内容
		copy(p, c.buffer[c.read:])

		// 已读取的长度
		readToEnd := c.size - c.read

		// p 中可拷贝的剩余空间
		left := canReadLen - readToEnd

		copy(p[readToEnd:], c.buffer[0:left])

		// start|__read_write__________|end
		c.read = left
	}

	c.isFull = false
	return canReadLen, nil
}

// ReadN 读取长度 n 的内容到 p 中
// 只存在读取了长度 n 的内容到 p 中，和 一点都不读取两种情况
// 不存在只读了部分(< n) 的内容到 p 中的情况
func (c *Circle) ReadN(n int, p []byte) error {
	if n == 0 || p == nil || len(p) < n {
		return ErrInvalidBuffer
	}

	if c.Len() < n {
		return ErrNotEnough
	}

	_, err := c.Read(p[:n])
	return err
}

// Peek 读取 n 个长度的内容，但不会产生任何影响
// 不会改变 c.read 和 c.write
func (c *Circle) Peek(n int, p []byte) error {
	write, read := c.write, c.read

	err := c.ReadN(n, p)
	c.write, c.read = write, read

	return err
}

// Skip 跳过 n 个字段，会改变 c.read
func (c *Circle) Skip(n int) error {
	if n <= 0 {
		return nil
	}

	if c.IsEmpty() {
		return ErrIsEmpty
	}

	// start|___read______write____|end
	if c.write > c.read {
		if c.read+n > c.write {
			return ErrInvalidSkipSize
		}

		c.read += n
	} else {
		// c.write <= c.read
		// start|___write______read____|end
		if c.read+n <= c.size {
			c.read += n
		} else {
			left := n - (c.size - c.read)
			if left > c.write {
				return ErrInvalidSkipSize
			}

			c.read = left
		}
	}

	c.isFull = false
	return nil
}

// IsEmpty 缓冲区是否为空
func (c *Circle) IsEmpty() bool {
	return !c.isFull && c.write == c.read
}

// IsFull 缓冲区是否已满
func (c *Circle) IsFull() bool {
	return c.isFull
}

func (c *Circle) String() string {
	if c.write >= c.read {
		return fmt.Sprintf("read: %d, write: %d, size: %d, len: %d, free: %d, isFull: %t, buffer: %v",
			c.read, c.write, c.size, c.Len(), c.Free(), c.isFull, c.buffer,
		)
	}

	return fmt.Sprintf("write: %d, read: %d, size: %d, len: %d, free: %d, isFull: %t, buffer: %v",
		c.write, c.read, c.size, c.Len(), c.Free(), c.isFull, c.buffer,
	)
}
