package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var ret int
	for i:=0; i<len(nums)-1;i=i+2 {
		ret += nums[i]
	}
	return ret
}

func main() {

}
