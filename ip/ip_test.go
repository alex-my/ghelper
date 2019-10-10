package ip_test

import (
	"testing"

	"github.com/alex-my/ghelper/ip"
)

func TestIPToInt(t *testing.T) {
	n := ip.ToUint32("192.168.1.1")
	if n != 3232235777 {
		t.Fatalf("n must be 3232235777, now: %d", n)
	}
}

func TestIPToString(t *testing.T) {
	s := ip.ToString(3232235777)
	if s != "192.168.1.1" {
		t.Fatalf("s must be 192.168.1.1, now: %s", s)
	}
}
