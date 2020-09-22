package random_test

import (
	"testing"

	"github.com/alex-my/ghelper/random"
)

func TestRandom(t *testing.T) {
	size := 10
	r1 := random.String(size)
	r2 := random.String(size)
	t.Logf("r1: %s", r1)
	t.Logf("r2: %s", r2)
	if len(r1) != size {
		t.Errorf("R1 length: %d is not size: %d", len(r1), size)
	}

	if len(r2) != size {
		t.Errorf("R2 length: %d is not size: %d", len(r2), size)
	}
}
