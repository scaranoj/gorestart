package main

import "testing"

type testcase struct {
	data   []int
	answer int
}

var testcases = []testcase{
	{[]int{21, 21}, 42},
	{[]int{11, 11}, 22},
	{[]int{-1, 0, 1}, 0},
	{[]int{350, 6}, 356},
}

//create a for to range over your slice elements

func TestMySum(t *testing.T) {
	for _, tc := range testcases {
		v := mySum(tc.data...)
		if v != tc.answer {
			t.Error("Expected ", tc.answer, "Got ", v)
		}
	}
}
