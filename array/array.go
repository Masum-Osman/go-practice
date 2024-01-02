package array

import (
	"fmt"
)

func Sum(numbers []int) int {
	sum := 0
	/*
		for i := 0; i < 5; i++ {
			sum += numbers[i]
		}
	*/

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func ContainsDuplicateTwoPointer(nums []int) bool {

	start := 0
	end := len(nums) - 1

	for start <= end {
		// fmt.Println(nums[start], nums[end])
		fmt.Println(start, end)
		if nums[start] == nums[end] && len(nums) != 1 && start != end {
			return true
		} else {
			start++
			end--
		}
	}

	return false
}

func ContainsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func ContainsDuplicateUsingMap(nums []int) bool {
	nums_map := make(map[int]int)
	nums_map[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		val, exist := nums_map[nums[i]]
		fmt.Println(val, exist)
	}

	// nums_map := map[int]int{}

	// for _, n := range nums {
	// 	if _, ok := nums_map[n]; !ok {
	// 		nums_map[n] = 1
	// 	} else {
	// 		return true
	// 	}
	// }
	// return false
	return false
}

func TwoDArray() {
	a := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
	}

	fmt.Println(a)

	c := [5][5]uint8{}
	fmt.Println(c)
}

func Travers2DArray() {
	/*
		So a = [5] means 5 is the row number.
	*/

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func LinearSearch2DArray(arr [][]int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if target == arr[i][j] {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func BinarySearch2DArray(arr [][]int, target int) []int {
	// O(n+m)
	var row int = 0
	var col int = len(arr[row]) - 1

	for row < len(arr) && col >= 0 {
		if arr[row][col] == target {
			return []int{row, col}
		}

		if arr[row][col] < target {
			row++
		} else {
			col--
		}
	}

	return []int{-1, -1}
}

func BinarySearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	low, high := 0, rows*cols-1

	for low <= high {
		mid := low + (high-low)/2
		midValue := matrix[mid/cols][mid%cols]

		if midValue == target {
			return true
		} else if midValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func BinarySearchMatrixFromNet(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	i := 0

	for i < rows {
		if target <= matrix[i][cols-1] {
			j := 0
			for j < cols {
				if target == matrix[i][j] {
					return true
				} else if target < matrix[i][j] {
					return false
				} else {
					j++
				}
			}
			return false
		}
		i++
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	begin := 0
	end := m*n - 1

	for begin <= end {
		mid := begin + (end-begin+1)/2
		val := matrix[mid/n][mid%n]

		if val == target {
			return true
		} else if val > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return false
}
