/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]. */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		if j, ok := tmpMap[target-nums[i]]; ok {
			return []int{j, i}
		}
		tmpMap[nums[i]] = i
	}
	return nil
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("%v", twoSum(arr, target))
}
