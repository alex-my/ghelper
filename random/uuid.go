package random

import (
	"encoding/binary"
	"encoding/hex"
	"sync/atomic"
	"time"

	"github.com/alex-my/ghelper/os"
)

var (
	machineID []byte
	processID []byte
	counter   uint32
)

// NewUUID 全局唯一标识符，12byte
func NewUUID() string {
	b := make([]byte, 12)
	// 1-4: Unix 时间戳
	binary.BigEndian.PutUint32(b, uint32(time.Now().Unix()))
	// 5-7: 主机名
	b[4], b[5], b[6] = machineID[0], machineID[1], machineID[2]
	// 8-9: 进程ID
	b[7], b[8] = processID[0], processID[1]
	// 递增
	counter := atomic.AddUint32(&counter, 1)
	b[9], b[10], b[11] = byte(counter>>16), byte(counter>>8), byte(counter)

	return hex.EncodeToString(b)
}

func init() {
	machineID = os.MachineID()
	processID = os.ProcessID()
	counter = Uint32()
}
