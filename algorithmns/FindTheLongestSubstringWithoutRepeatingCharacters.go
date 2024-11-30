package main

import "fmt"

// Fuction to find the length longest substring without repeating characters

func lengthOfTheLongestSubstring(s string) int {

	// The make built-in function allocates and initializes an object of type
	// slice, map, or chan (only). Like new, the first argument is a type, not a
	// value. Unlike new, make's return type is the same as the type of its
	// argument, not a pointer to it. The specification of the result depends on
	// the type:
	//
	//	Slice: The size specifies the length. The capacity of the slice is
	//	equal to its length. A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length. For example, make([]int, 0, 10) allocates an underlying array
	//	of size 10 and returns a slice of length 0 and capacity 10 that is
	//	backed by this underlying array.
	//	Map: An empty map is allocated with enough space to hold the
	//	specified number of elements. The size may be omitted, in which case
	//	a small starting size is allocated.
	//	Channel: The channel's buffer is initialized with the specified
	//	buffer capacity. If zero, or the size is omitted, the channel is
	//	unbuffered.
	charMap := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {

		if pos, found := charMap[s[right]]; found && pos >= left {
			left = pos + 1
		}

		// pos, found := charMap[s[right]]:

		// charMap[s[right]] checks if the current character (s[right]) has already been encountered and retrieves its last recorded position (pos).
		// pos: The last index where the character s[right] was seen.
		// found: A boolean value indicating whether the character s[right] was found in the charMap.
		// if found && pos >= left:

		// This is a conditional check combining two conditions:
		// found:
		// Checks if the character s[right] is already in the map (found is true if the character exists in charMap).
		// pos >= left:
		// pos is the previous position of the character.
		// left is the starting index of the current substring.
		// If pos >= left, it means the character is within the current substring. This indicates a repeating character within the current sliding window.
		// left = pos + 1:

		// When a repeating character is found in the current substring:
		// Move the left pointer just past the last occurrence of the character.
		// This step effectively excludes the repeating character from the new substring window.
		// This adjustment ensures that the new substring starting from left contains only unique characters.

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
