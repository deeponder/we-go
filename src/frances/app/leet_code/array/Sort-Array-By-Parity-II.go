package main

func sortArrayByParityII(A []int) []int {
	i := 1
	for j := 0; j<len(A)-1; j=j+2 {
		if A[j] % 2 == 1 {
			for A[i] % 2 == 1 {
				i = i+2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

func main() {
	
}
