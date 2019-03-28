package main

//动态规划 dp方程
func minCostClimbingStairs(cost []int) int {
	lenCost := len(cost)
	dp := make([]int, lenCost)
	copy(dp, cost)
	for i:=2; i<len(cost); i++ {
		dp[i] += miniCost(dp[i-1], dp[i-2])
	}

	return miniCost(dp[lenCost-1], dp[lenCost-2])
}

func miniCost(a int, b int) int {
	if a > b {
		return b
	}	 else {
		return a
	}
}

func main() {
	
}
