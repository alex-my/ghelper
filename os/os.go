package os

import (
	"crypto/md5"
	"crypto/rand"
	"io"
	"os"
)

// MachineID 获取机器标识符，如果获取不到则随机
func MachineID() []byte {
	id := make([]byte, 3)

	if hostname, err := os.Hostname(); err == nil {
		m := md5.New()
		m.Write([]byte(hostname))
		// 只复制前3个
		// [10 93 205 142 83 34 228 215 53 108 250 67 95 45 218 199]
		// [10 93 205]
		copy(id, m.Sum(nil))
		return id
	}

	if _, err := io.ReadFull(rand.Reader, id); err == nil {
		return id
	}
	panic("Failed to get hostname")
}

// ProcessID 获取进程信息
func ProcessID() []byte {
	pid := os.Getpid()

	b := make([]byte, 2)
	b[0] = byte(pid >> 8)
	b[1] = byte(pid)
	return b
}
