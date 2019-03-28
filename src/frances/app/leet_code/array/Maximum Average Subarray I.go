package main


func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i:=0; i<k; i++ {
		max += nums[i]
	}

	for j:=k; j<len(nums); j++ {
		tmp := 0
		for n:=j; n>j-k; n-- {
			tmp += nums[n]
		}
		if tmp>max{
			max = tmp
		}
	}

	return float64(max)/float64(k)
}

func main() {
	
}
