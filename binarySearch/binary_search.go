package binarySearch

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

func FirstBadVersion(n int) int {
	left := 0
	right := n
	lastBadVersion := 1

	for i := left; i < right; i++ {

	}

	return lastBadVersion
}
