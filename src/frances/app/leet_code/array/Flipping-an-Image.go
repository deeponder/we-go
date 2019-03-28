package main

func flipAndInvertImage(A [][]int) [][]int {
	for i:=0; i< len(A); i++ {
		last := len(A[i])-1
		for j:=0; j<=last/2; j++ {
			A[i][j], A[i][last-j] = A[i][last-j] ^ 1, A[i][j] ^ 1
		}
	}
	return A
}

func main() {
	
}
