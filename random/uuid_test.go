package random_test

import (
	"testing"

	"github.com/alex-my/ghelper/random"
)

func TestUUID(t *testing.T) {
	id1 := random.NewUUID()
	id2 := random.NewUUID()
	t.Logf("id1: %s\n", id1)
	t.Logf("id2: %s\n", id2)
	if id1 == id2 {
		t.Error("Id1 and id2 cannot be the same")
	}
}
