package main

func findMaxConsecutiveOnes(nums []int) int {
	maxnum := 0
	tmp := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != 1 {
			tmp = 0
		}

		if nums[i] == 1 {
			tmp ++
		}

		if tmp > maxnum {
			maxnum = tmp
		}
	}

	return maxnum
}

func main() {
	
}
