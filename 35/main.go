package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	// Ví dụ 1
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(searchInsert(nums1, target1)) // Output: 2

	// Ví dụ 2
	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Println(searchInsert(nums2, target2)) // Output: 1

	// Ví dụ 3
	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	fmt.Println(searchInsert(nums3, target3)) // Output: 4
}

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

// Constraints:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums contains distinct values sorted in ascending order.
// -104 <= target <= 104
