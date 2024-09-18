package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicates
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // skip duplicates
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++ // skip duplicates
					}
					for left < right && nums[right] == nums[right-1] {
						right-- // skip duplicates
					}
					left++
					right--
				}
			}
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 0, -1, 0, -2, 2}
	target1 := 0
	fmt.Println(fourSum(nums1, target1)) // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println(fourSum(nums2, target2)) // Output: [[2,2,2,2]]
}

// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.

// Example 1:

// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// Example 2:

// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]

// Constraints:

// 1 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
