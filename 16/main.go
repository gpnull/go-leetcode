package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // Sắp xếp mảng để sử dụng hai con trỏ
	closest := math.MaxInt32
	n := len(nums)

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1 // Khởi tạo hai con trỏ
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum < target {
				left++ // Tăng con trỏ trái
			} else if sum > target {
				right-- // Giảm con trỏ phải
			} else {
				return sum // Nếu sum bằng target, trả về ngay
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // Output: 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // Output: 0
}

// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

// Return the sum of the three integers.

// You may assume that each input would have exactly one solution.

// Example 1:

// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// Example 2:

// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

// Constraints:

// 3 <= nums.length <= 500
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
