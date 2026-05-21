package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(2, 2)
	if total != 4 {
		t.Errorf("Got %d, want %d", total, 4)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(4, 2)
	if total != 2 {
		t.Errorf("Got %d, want %d", total, 2)
	}
}

func TestDoMathAddTest(t *testing.T) {
	totalAdd := doMath(10, 10, add)
	if totalAdd != 20 {
		t.Errorf("Got %d, want %d", totalAdd, 20)
	}
}

func TestDoMathSubTest(t *testing.T) {
	totalSub := doMath(10, 10, subtract)
	if totalSub != 0 {
		t.Errorf("Got %d, want %d", totalSub, 0)
	}
}
