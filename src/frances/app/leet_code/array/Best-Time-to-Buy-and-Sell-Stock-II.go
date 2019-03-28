package main

//短线操作， 后一天大于前一天就卖出
func maxProfit(prices []int) int {
	max := 0
	for i:=1; i<len(prices);i++ {
		if prices[i] - prices[i-1] > 0 {
			max += prices[i]-prices[i-1]
		}
	}
	return max
}

func main() {
	println(maxProfit([]int{1,2,3,4,5}))
}
