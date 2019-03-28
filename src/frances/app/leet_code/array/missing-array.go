package main

func missingNumber(nums []int) int {
	numsLen := len(nums)
	total := (0+numsLen) * (numsLen+1) / 2
	inputTotal := 0
	for _, v := range nums {
		inputTotal += v
	}

	return total - inputTotal
}

func main() {
	
}
