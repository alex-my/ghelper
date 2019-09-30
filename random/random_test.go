package random

import (
	"testing"
)

func TestRandom(t *testing.T) {
	size := 10
	r1 := String(size)
	r2 := String(size)
	t.Logf("r1: %s", r1)
	t.Logf("r2: %s", r2)
	if r1 == r2 {
		t.Error("R1 and r2 cannot be the same")
	}
	if len(r1) != size {
		t.Errorf("R1 length: %d is not size: %d", len(r1), size)
	}

	if len(r2) != size {
		t.Errorf("R2 length: %d is not size: %d", len(r2), size)
	}
}
