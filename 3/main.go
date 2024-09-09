package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if index, found := charIndexMap[s[i]]; found && index >= start {
			start = index + 1
		}
		charIndexMap[s[i]] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	examples := []string{"abcabcbb", "bbbbb", "pwwkew"}
	for _, example := range examples {
		fmt.Printf("Input: %s, Output: %d\n", example, lengthOfLongestSubstring(example))
	}
}

// Given a string s, find the length of the longest
// substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
