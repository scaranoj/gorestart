package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	v := 0
	v = mySum(5, 5)
	if v != 10 {
		t.Error("Expected 10, got ", v)
	}

}
