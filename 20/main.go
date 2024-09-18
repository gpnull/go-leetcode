package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if open, ok := mapping[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			stack = append(stack, char) // push
		}
	}
	return len(stack) == 0
}

func main() {
	examples := []string{"()", "()[]{}", "(]", "([])"}
	for _, example := range examples {
		fmt.Printf("Input: %s, Output: %v\n", example, isValid(example))
	}
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

// Constraints:

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.
