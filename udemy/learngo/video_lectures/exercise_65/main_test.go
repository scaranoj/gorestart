package main

import (
	"testing"
)

func MultiplyTest(t *testing.T) {
	got := 5 * 5
	want := 10
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
