package main

import (
	"fmt"
)

var digitToLetters = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var combinations []string
	var backtrack func(index int, path string)

	backtrack = func(index int, path string) {
		if index == len(digits) {
			combinations = append(combinations, path)
			return
		}

		letters := digitToLetters[rune(digits[index])]
		for _, letter := range letters {
			backtrack(index+1, path+string(letter))
		}
	}

	backtrack(0, "")
	return combinations
}

func main() {
	examples := []string{"23", "", "2"}
	for _, example := range examples {
		fmt.Printf("Input: %s, Output: %v\n", example, letterCombinations(example))
	}
}

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:

// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:

// Input: digits = ""
// Output: []
// Example 3:

// Input: digits = "2"
// Output: ["a","b","c"]

// Constraints:

// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
