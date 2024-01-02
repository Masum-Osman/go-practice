package array

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	/*
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	*/

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestContainsDuplicates(t *testing.T) {
	type testCases struct {
		input  []int
		result bool
	}

	cases := []testCases{
		{[]int{7, 5, -2, -4, 0}, false},
		// {[]int{1, 2, 3, 1}, true},
		// {[]int{0, 4, 5, 0, 3, 6}, true},
		// {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range cases {
		got := ContainsDuplicateUsingMap(tc.input)

		if tc.result != got {
			t.Error("Wanted: ", tc.result, " || Got: ", got)
		}
	}
}

func TestTwoDArray(t *testing.T) {
	Travers2DArray()
}

func TestLinearSearchIn2DArray(t *testing.T) {

	type testCases struct {
		nums   [][]int
		target int
		result bool
	}

	cases := []testCases{
		// {[][]int{{3, 12, 9}, {5, 2, 89}, {90, 45, 22}}, 89, true},
		// {[][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}, 2, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 5, true},
		// {[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	for _, tc := range cases {
		got := BinarySearchMatrix(tc.nums, tc.target)

		if tc.result != got {
			t.Error("Wanted: ", tc.result, " || Got: ", got)
		}

		fmt.Println(got)
	}
}
