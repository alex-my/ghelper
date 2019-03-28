package array

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	il := []int{1, 2, 3, 4}
	dst := "1,2,3,4"
	if IntToString(il) != dst {
		t.Error("IntToString err")
	}
}

func TestInt64ToString(t *testing.T) {
	il := []int64{1, 2, 3, 4}
	dst := "1,2,3,4"
	if Int64ToString(il) != dst {
		t.Error("Int64ToString err")
	}
}

func TestStringToString(t *testing.T) {
	il := []string{"1", "2", "3", "4"}
	dst := "1,2,3,4"
	if StringToString(il) != dst {
		t.Error("StringToString err")
	}
}
