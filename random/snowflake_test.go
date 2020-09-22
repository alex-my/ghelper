package random_test

import (
	"sync"
	"testing"

	"github.com/alex-my/ghelper/random"
)

func TestSnowflake(t *testing.T) {
	// 多个协程并发生成不重复的 uuid

	workerNum := 5
	uuidNum := 10000

	var wg sync.WaitGroup
	wg.Add(workerNum)

	ch := make(chan uint64, workerNum*uuidNum+1)

	for i := 0; i < workerNum; i++ {
		go func(workderID uint32) {
			defer wg.Done()

			snowflake, _ := random.NewSnowflake(workderID)

			for j := 0; j < uuidNum; j++ {
				uuid, _ := snowflake.UUID()
				ch <- uuid
			}

		}(uint32(i))
	}

	wg.Wait()

	// 检查是否有重复
	m := make(map[uint64]int, workerNum*uuidNum)

	for i := 0; i < workerNum*uuidNum; i++ {
		uuid := <-ch
		if _, exist := m[uuid]; exist {
			t.Fatalf("%d repeated", uuid)
		}

		m[uuid] = 0
	}
}
