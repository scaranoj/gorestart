package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	if sum != 4 {
		t.Errorf("Want %d, got %d", 4, sum)
	}
}
