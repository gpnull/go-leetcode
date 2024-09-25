package main

import "fmt"

func longestValidParentheses(s string) int {
	maxLength := 0
	stack := []int{-1} // khởi tạo ngăn xếp với -1 để xử lý các trường hợp hợp lệ

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i) // thêm chỉ số vào ngăn xếp
		} else {
			stack = stack[:len(stack)-1] // xóa phần tử trên cùng
			if len(stack) == 0 {
				stack = append(stack, i) // thêm chỉ số ')' vào ngăn xếp
			} else {
				maxLength = max(maxLength, i-stack[len(stack)-1]) // cập nhật độ dài tối đa
			}
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	examples := []string{"(()", ")()())", "", ")()("}
	for _, example := range examples {
		fmt.Printf("Input: %s, Output: %d\n", example, longestValidParentheses(example))
	}
}

// Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
// substring
// .

// Example 1:

// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// Example 2:

// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// Example 3:

// Input: s = ""
// Output: 0

// Constraints:

// 0 <= s.length <= 3 * 104
// s[i] is '(', or ')'.
