package main

import "testing"

func TestMySum(t *testing.T) {
	v := mySum(2, 2)
	if v != 4 {
		t.Error("Expected 2, got", v)
	}
}
