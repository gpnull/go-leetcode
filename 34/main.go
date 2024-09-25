package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	left, right := -1, -1
	// Tìm vị trí bắt đầu
	for l, r := 0, len(nums)-1; l <= r; {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		left = l
	}

	// Tìm vị trí kết thúc
	for l, r := 0, len(nums)-1; l <= r; {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		right = r
	}

	if left <= right {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func main() {
	// Ví dụ 1
	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	fmt.Println(searchRange(nums1, target1)) // Output: [3, 4]

	// Ví dụ 2
	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 6
	fmt.Println(searchRange(nums2, target2)) // Output: [-1, -1]

	// Ví dụ 3
	nums3 := []int{}
	target3 := 0
	fmt.Println(searchRange(nums3, target3)) // Output: [-1, -1]
}

// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums is a non-decreasing array.
// -109 <= target <= 109
