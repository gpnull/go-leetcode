package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	indices := []int{}

	for i := 0; i < wordLen; i++ { // Lặp qua từng vị trí bắt đầu
		left := i
		count := 0
		seen := make(map[string]int)

		for right := i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if _, ok := wordMap[word]; ok {
				seen[word]++
				count++

				// Nếu từ đã xuất hiện quá số lần cho phép, dịch cửa sổ
				for seen[word] > wordMap[word] {
					seen[s[left:left+wordLen]]--
					count--
					left += wordLen
				}

				// Nếu số từ đã thấy bằng số từ cần tìm, lưu chỉ số
				if count == wordCount {
					indices = append(indices, left)
				}
			} else {
				// Reset nếu từ không hợp lệ
				seen = make(map[string]int)
				count = 0
				left = right + wordLen
			}
		}
	}

	return indices
}

func main() {
	examples := []struct {
		s     string
		words []string
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}},
	}

	for _, example := range examples {
		result := findSubstring(example.s, example.words)
		fmt.Printf("Input: s = %q, words = %v\nOutput: %v\n", example.s, example.words, result)
	}
}

// You are given a string s and an array of strings words. All the strings of words are of the same length.

// A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

// For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
// Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

// Example 1:

// Input: s = "barfoothefoobarman", words = ["foo","bar"]

// Output: [0,9]

// Explanation:

// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

// Example 2:

// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

// Output: []

// Explanation:

// There is no concatenated substring.

// Example 3:

// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

// Output: [6,9,12]

// Explanation:

// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

// Constraints:

// 1 <= s.length <= 104
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// s and words[i] consist of lowercase English letters.
