package main

//同股票的最大收益
//动态规划， dp方程 dp[i+1] = max(dp[i]+dp[i]+a[i])
func maxSubArray(nums []int) int {
	max := nums[0]
	tmp := 0
	for _,v := range nums {
		tmp = tmp + v
		if tmp > max {
			max = tmp
		}
		//重置
		if tmp<0 {
			tmp = 0
		}
	}
	return max
}

func main() {
	maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
}
