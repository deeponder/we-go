/*Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit. Example 1:

Input: [7, 1, 5, 3, 6, 4]
Output: 5
max. difference = 6-1 = 5 (not 7-1 = 6, as selling price needs to be larger than buying price)

Example 2:

Input: [7, 6, 4, 3, 1]
Output: 0
In this case, no transaction is done, i.e. max profit = 0. */
package main

//1. 只需要关心总的收益最大， 不需要关心哪天卖出
//2. 遍历时，记录这么多天下来的累积收益即可， 大于max, 则替换
//3. 动态规划问题， dp方程， dp[i+1] = max(dp[i], dp[i]+a[i])
func maxProfit(prices []int) int {
	max := 0
	tmp := 0
	for i:=1; i<len(prices); i++ {
		tmp = tmp + (prices[i] - prices[i-1])
		if tmp < 0 {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(input)
	println(maxProfit)
}
