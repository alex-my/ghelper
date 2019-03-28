package random

import (
	"testing"
)

func TestUUID(t *testing.T) {
	id1 := NewUUID()
	id2 := NewUUID()
	t.Logf("id1: %s\n", id1)
	t.Logf("id2: %s\n", id2)
	if id1 == id2 {
		t.Error("Id1 and id2 cannot be the same")
	}
}
