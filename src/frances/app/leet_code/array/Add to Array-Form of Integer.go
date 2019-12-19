package main

import (
	"fmt"
)



func addToArrayForm(A []int, K int) []int {
	//小学加法运算
	//ret := make([]int, 0)
	//cur := K
	//i := len(A)-1
	//for i>=0 || cur>0{
	//	if i>=0 {
	//		cur = cur + A[i]
	//	}
	//
	//	ret = append(ret, cur%10)
	//	cur = cur/10
	//	i--
	//}
	//
	//result := make([]int, 0)
	//for i:=len(ret)-1; i>=0; i--{
	//	result = append(result, ret[i])
	//}
	//return result

	// !!!! wrong, out of range!!!
	arrNum := 0
	lenA := len(A)
	for i:=0; i<lenA; i++{
		tmp:=A[i]
		for j:=0; j<lenA-i-1; j++{
			tmp = tmp*10
		}
		arrNum = arrNum+tmp
	}
	resultNum := arrNum + K
	resultRevers := make([]int, 0)
	for {
		resultRevers = append(resultRevers, resultNum%10)
		if resultNum/10 == 0 {
			break
		}
		resultNum = resultNum/10
	}

	result := make([]int, 0)
	for i:=len(resultRevers)-1; i>=0;i--{
		result = append(result, resultRevers[i])
	}

	return result
}

func main() {
	fmt.Printf("%+v", addToArrayForm([]int{1,2,6,3,0,7,1,7,1,9,7,5,6,6,4,4,0,0,6,3}, 21))
}
