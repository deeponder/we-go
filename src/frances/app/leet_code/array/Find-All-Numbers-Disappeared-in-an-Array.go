package main

import "sort"

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	//sort.Ints(nums)
	for i:=0;i<len(nums); i++ {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for i:=0; i<len(nums); i++ {
		if nums[i] != i+1 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	
}
