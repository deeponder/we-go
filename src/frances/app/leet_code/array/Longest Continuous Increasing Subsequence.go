package main

func findLengthOfLCIS(nums []int) int {
	var max int
	if len(nums) == 0 {
		max = 0
	} else {
		max = 1
	}
	tmp := 1
	for i:=0; i<len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			tmp ++
		} else {
			tmp = 1
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	println(findLengthOfLCIS([]int{2,2,2,2,2}))
}
