package main

import "fmt"


//1、元素的交换  拷贝、加减
//2. go中数组和切片的区别
func rotate(nums []int, k int)  {
	numLen := len(nums)
	tmpArr := make([]int, numLen)
	copy(tmpArr, nums)
	for i:= 0; i<numLen; i++ {
		nums[(k+i)%numLen] = tmpArr[i]
	}
}

func main() {
	//rotate([]int{1,2,3,4,5,6,7},3)
	num := [3]int{1,2,3}
	tmpArr := num
	tmpArr[1] = 3
	fmt.Printf("%+v\n %+v", tmpArr, num)
}
