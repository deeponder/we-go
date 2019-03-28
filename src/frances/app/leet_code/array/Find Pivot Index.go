package main


func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if len(nums) == 0 {
		return -1
	}
	if sum-nums[0] == 0 {
		return 0
	}

	halfSum := 0
	for i:=0; i<len(nums)-1; i++{
		halfSum += nums[i]
		if halfSum == (sum-nums[i+1])/2 && (sum-nums[i+1])%2==0{
			return i+1
		}
	}

	return -1
}

func main() {
	
}
