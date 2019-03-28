package main

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	lastProduct := nums[numsLen-1]*nums[numsLen-2]*nums[numsLen-3]
	firstProduct := nums[0]*nums[1]*nums[numsLen-1]
	if lastProduct > firstProduct {
		return lastProduct
	} else {
		return firstProduct
	}
}

func main() {
	
}
