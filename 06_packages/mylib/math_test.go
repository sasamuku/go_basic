package mylib

import "testing"

func TestAverage(t *testing.T) { // 関数名：TestXxx
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
