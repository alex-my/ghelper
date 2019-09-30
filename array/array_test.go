package array

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	if IntToString([]int{1, 2, 3, 4}) != "1,2,3,4" {
		t.Error("IntToString err")
	}

	if IntToString([]int{1, 2, 3, 4}, "+") != "1+2+3+4" {
		t.Error("IntToString err, +")
	}
}

func TestInt64ToString(t *testing.T) {
	if Int64ToString([]int64{1, 2, 3, 4}) != "1,2,3,4" {
		t.Error("Int64ToString err")
	}

	if Int64ToString([]int64{1, 2, 3, 4}, "+") != "1+2+3+4" {
		t.Error("Int64ToString err, +")
	}
}

func TestStringToString(t *testing.T) {
	if StringToString([]string{"1", "2", "3", "4"}) != "1,2,3,4" {
		t.Error("StringToString err")
	}

	if StringToString([]string{"1", "2", "3", "4"}, "+") != "1+2+3+4" {
		t.Error("StringToString err, +")
	}
}
