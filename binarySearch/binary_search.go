package binarySearch

import "fmt"

func Search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for i := left; i < right; i++ {
		if target == nums[i] {
			return i
		} else if target > nums[i] {
			left = (left + right) / 2
		} else {
			right = (left + right) / 2
		}
	}

	return -1
}

func SearchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for i := left; i < right; i++ {
		mid := (left + right) / 2
		fmt.Println("LEFT: ", left, "RIGHT: ", right, "\t", "MID: ", mid, "\t",
			"NUMS[MIS]: ", nums[mid])

		if nums[mid] >= target {
			right = mid
			// return mid
		} else {
			left = mid + 1
		}
	}
	return left
}
