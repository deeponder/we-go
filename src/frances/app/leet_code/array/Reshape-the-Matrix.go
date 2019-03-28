package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	m := len(nums[0])

	if n * m != r * c {
		return nums
	} else {
		oneArr := make([]int,0)
		retArr := make([][]int, r)
		for _, v := range nums {
			for _, num := range v {
				oneArr = append(oneArr, num)
			}
		}
		//fmt.Printf("%+v", oneArr)
		var z= 0
		for j := range retArr {
			retArr[j] = make([]int, c)
			for k := 0; k < c; k++ {
				retArr[j][k] = oneArr[z]
				z++
			}
		}
		return retArr
	}
}

func main() {
	fmt.Printf("%+v", matrixReshape([][]int{{1,2}, {3,4}}, 1, 4))
}
