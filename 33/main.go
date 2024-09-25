package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Determine which side is sorted
		if nums[left] <= nums[mid] { // Left side is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Target is in the left side
			} else {
				left = mid + 1 // Target is in the right side
			}
		} else { // Right side is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Target is in the right side
			} else {
				right = mid - 1 // Target is in the left side
			}
		}
	}
	return -1 // Target not found
}

func main() {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	fmt.Println(search(nums1, target1)) // Output: 4

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	fmt.Println(search(nums2, target2)) // Output: -1

	nums3 := []int{1}
	target3 := 0
	fmt.Println(search(nums3, target3)) // Output: -1
}

// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
// Example 3:

// Input: nums = [1], target = 0
// Output: -1

// Constraints:

// 1 <= nums.length <= 5000
// -104 <= nums[i] <= 104
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -104 <= target <= 104
