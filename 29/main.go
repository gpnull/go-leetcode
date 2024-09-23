package main

import "fmt"

func divide(dividend int, divisor int) int {
	// xử lý trường hợp tràn số
	if dividend == -1<<31 && divisor == -1 {
		return (1 << 31) - 1 // Trả về 231 - 1
	}

	// xác định dấu của kết quả
	negative := (dividend < 0) != (divisor < 0)

	// chuyển đổi sang số dương
	dividend, divisor = abs(dividend), abs(divisor)
	quotient := 0

	// thực hiện phép chia bằng cách dịch bit
	for dividend >= divisor {
		temp, multiple := divisor, 1
		for dividend >= temp {
			dividend -= temp
			quotient += multiple

			// dịch bit để tăng tốc độ
			temp <<= 1
			multiple <<= 1
		}
	}

	// trả về kết quả với dấu
	if negative {
		return -quotient
	}
	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(-8, 2))
	fmt.Println(divide(-8, -2))
	fmt.Println(divide(2147483647, -1))
}

// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

// The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

// Return the quotient after dividing dividend by divisor.

// Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.

// Example 1:

// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = 3.33333.. which is truncated to 3.
// Example 2:

// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = -2.33333.. which is truncated to -2.

// Constraints:

// -231 <= dividend, divisor <= 231 - 1
// divisor != 0
