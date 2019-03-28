package main

import "fmt"

func findShortestSubArray(nums []int) int {
	//存元素的开始位置， 计算距离
	elePos := make(map[int]int, 0)
	//存元素出现的个数， 判断是否替换最短距离
	eleCount := make(map[int]int, 0)
	shortestLen := 0
	maxCount := 0
	for i:=0;i<len(nums);i++ {
		if _, ok := elePos[nums[i]]; ok {
			eleCount[nums[i]]++
			if eleCount[nums[i]] > maxCount {
				maxCount = eleCount[nums[i]]
				shortestLen = i - elePos[nums[i]] + 1
			} else if eleCount[nums[i]] == maxCount {
				if shortestLen > i - elePos[nums[i]] + 1 {
					shortestLen = i - elePos[nums[i]] + 1
				}
			}
		} else {
			elePos[nums[i]] = i
		}

	}
	if shortestLen == 0 {
		shortestLen = 1
	}

	return shortestLen
}

func main() {
	println(findShortestSubArray([]int{1}))
}
