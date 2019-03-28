package main

func getRow(rowIndex int) []int {
	//初始化化数组长度
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i:=0; i < rowIndex; i++ {
		res = append(res, 1)
		//从后往前赋值， 滚动计算
		for j := i; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}
	}

	return res
}

func main() {
	
}
