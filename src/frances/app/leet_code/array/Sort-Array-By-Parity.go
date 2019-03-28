package main

import "fmt"

func sortArrayByParity(A []int) []int {
	tmp := make([]int, len(A))
	j := len(A) - 1
	k := 0
	for i:=0; i<len(A); i++ {
		if A[i]%2 != 0 {
			tmp[j] = A[i]
			j--
		} else {
			tmp[k] = A[i]
			k++
		}
	}
	return tmp
}

func main() {
	fmt.Printf("%+v", sortArrayByParity([]int{3,1,2,4}))
}
