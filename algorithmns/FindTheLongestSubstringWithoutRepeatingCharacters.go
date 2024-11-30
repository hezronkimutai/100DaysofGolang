package main

import "fmt"

// Fuction to find the length longest substring without repeating characters

func lengthOfTheLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if pos, found := charMap[s[right]]; found && pos >= left {
			left = pos + 1
		}
		charMap[s[right]] = right

		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Package Declaration: The main package is required to indicate that this file is an executable program.
// Go programs start execution from the main function within the main package.
func main() {
	fmt.Println(lengthOfTheLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfTheLongestSubstring("bbbbbb"))
	fmt.Println(lengthOfTheLongestSubstring("hezron"))
	fmt.Println(lengthOfTheLongestSubstring("chelimo"))
	fmt.Println(lengthOfTheLongestSubstring("Kimmdh"))

}
