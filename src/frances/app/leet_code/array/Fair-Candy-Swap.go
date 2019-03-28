package main

//每次交换， 两个数组的和的差会变改变2*(A[i]-B[j])
func fairCandySwap(A []int, B []int) []int {
	suma := 0
	sumb := 0
	for _, v := range A {
		suma += v
	}

	for _, v := range B {
		sumb += v
	}

	diff := (suma - sumb)/2
	res := make([]int, 2)
	for i:=0; i<len(A); i++ {
		for j:=0; j<len(B); j++ {
			if A[i] == B[j] + diff {
				res[0] = A[i]
				res[1] = B[j]
			}
		}
	}
	return res
}

func main() {
	
}
