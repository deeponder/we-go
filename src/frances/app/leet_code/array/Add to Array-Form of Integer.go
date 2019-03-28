package main

import (
	"fmt"
)



func addToArrayForm(A []int, K int) []int {
	//小学加法运算
	ret := make([]int, 0)
	cur := K
	i := len(A)-1
	for i>=0 || cur>0{
		if i>=0 {
			cur = cur + A[i]
		}

		ret = append(ret, cur%10)
		cur = cur/10
		i--
	}

	result := make([]int, 0)
	for i:=len(ret)-1; i>=0; i--{
		result = append(result, ret[i])
	}
	return result
}

func main() {
	fmt.Printf("%+v", addToArrayForm([]int{2, 7, 4}, 21))
}
