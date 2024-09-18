package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(prefix) > len(str) || str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	// Example 1
	strs1 := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs1))

	// Example 2
	strs2 := []string{"rog", "racecar", "rar"}
	println(longestCommonPrefix(strs2))
}

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.
