// Here's another medium-difficulty algorithm challenge in Go:

// Problem: Two Sum
// Description: Given an array of integers nums and an integer target, return the indices of the two numbers such that they add up to target.
// You may assume that each input has exactly one solution, and you may not use the same element twice.

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		compliment := target - num
		if index, found := numMap[compliment]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 1, 15}
	nums1 := []int{3, 7, 1, 15, 6}
	nums2 := []int{2, 7, 1, 15}
	nums3 := []int{2, 8, 1, 15}

	target := 9
	result := twoSum(nums, target)
	result1 := twoSum(nums1, target)
	result2 := twoSum(nums2, target)
	result3 := twoSum(nums3, target)

	fmt.Println("array", nums, "Target", target, "Indices:", result)
	fmt.Println("array1", nums1, "Target1", target, "Indices1:", result1)
	fmt.Println("array2", nums2, "Target2", target, "Indices2:", result2)
	fmt.Println("array3", nums3, "Target3", target, "Indices3:", result3)

}

// Explanation:
// Hash Map Usage: A map (numMap) stores each number and its index as you iterate.
// Complement Check: For each number, compute the complement (target - num). Check if this complement exists in the map.
// Return Indices: If the complement is found, return the pair of indices.
// Complexity:
// Time: O(n), where n is the length of the input array.
// Space: O(n), to store the hash map.
