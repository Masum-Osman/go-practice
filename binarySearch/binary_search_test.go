package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	type testCases struct {
		nums   []int
		target int
		result int
	}

	cases := []testCases{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
	}

	for _, tc := range cases {
		got := Search(tc.nums, tc.target)

		t.Log(got)
		if tc.result != got {
			t.Error("Wanted: ", tc.result, "Got: ", got)
		}
	}
}
