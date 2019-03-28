package main

func containsNearbyDuplicate(nums []int, k int) bool {
	numPos := make(map[int]int, 0)
	for i:=0; i<len(nums); i++ {
		if v, ok := numPos[nums[i]]; ok {
			if i-v <= k {
				return true
			}
		}
		numPos[nums[i]] = i
	}

	return false
}

func main() {
	
}
