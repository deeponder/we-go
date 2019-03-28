/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
 */
package main

func isInSlice(nums []int, n int) bool {
	for _, v := range nums {
		if v == n {
			return true
		}
		if v > n {
			return false
		}
	}

	return false
}

func removeDuplicatesFromSortedArray(nums []int) int {
	j := 1
	for i:= 1; i<len(nums); i++ {
		if nums[i] > nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func main() {
	println(removeDuplicatesFromSortedArray([]int{1, 1, 2, 4, 2, 3, 1}))
}
