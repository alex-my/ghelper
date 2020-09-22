package random

import (
	"errors"
	"sync"

	"github.com/alex-my/ghelper/time"
)

// twitter 雪花算法
// 0 - 毫秒时间戳(41bit) - 机器 id(10bit) - 序列号(12bit)

var (
	// ErrWorkerID 无效的 workerID，取值范围 [0, maxWorkerID]
	ErrWorkerID = errors.New("bad worker id")
	// ErrTimeBackward 时间倒退，当前时间比上一次记录的时间还要小
	ErrTimeBackward = errors.New("time backward")
)

const (
	workerIDBits = 10

	sequenceBits = 12

	maxWorkerID = -1 ^ (-1 << workerIDBits)

	maxSequence = -1 ^ (-1 << sequenceBits)
)

// Snowflake 用来生成 uuid 的工具
type Snowflake struct {
	// 记录上一次产生 id 的毫秒时间戳
	lastTimestamp int64

	// 当前毫秒内已生成的序列号，从 0 开始, 0 - maxSequence
	sequence uint16

	// 用来表示不同节点，这样不同节点生成的一定不同， 0 - maxWorkerID
	workerID uint32

	lock sync.Mutex
}

// NewSnowflake ..
func NewSnowflake(workerID uint32) (*Snowflake, error) {
	if workerID > maxWorkerID {
		return nil, ErrWorkerID
	}

	return &Snowflake{
		workerID: workerID,
	}, nil
}

// UUID 获取 uuid
func (snowflake *Snowflake) UUID() (uint64, error) {
	snowflake.lock.Lock()
	defer snowflake.lock.Unlock()

	return snowflake.generateUUID()
}

func (snowflake *Snowflake) generateUUID() (uint64, error) {
	now := time.MS()
	if now < snowflake.lastTimestamp {
		return 0, ErrTimeBackward
	}

	if now == snowflake.lastTimestamp {
		// 上一次与当前都在同一个毫秒内，递增数量
		snowflake.sequence = (snowflake.sequence + 1) & maxSequence

		// 如果已经超出当前毫秒可以记录的范围 maxSequence
		// 1000 & 0111 => 0
		if snowflake.sequence == 0 {
			// 暂停到下一个毫秒
			for now == snowflake.lastTimestamp {
				time.SleepMircosecond(100)
				now = time.MS()
			}
		}
	} else {
		snowflake.sequence = 0
	}

	snowflake.lastTimestamp = now

	// 将相关数据封装成 uint64
	uuid := (uint64(snowflake.lastTimestamp) << (workerIDBits + sequenceBits)) | (uint64(snowflake.workerID) << sequenceBits) | (uint64(snowflake.sequence))
	return uuid, nil
}
