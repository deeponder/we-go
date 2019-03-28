package main

func transpose(A [][]int) [][]int {
	m := len(A[0])
	n := len(A)

	B := make([][]int, m)
	//for i:=0; i<m; i++ {
	//	B[i] = make([]int,n)
	//}
	for i:=0; i<m; i++ {
		B[i] = make([]int,n)
		for j:=0; j<n; j++ {
			B[i][j] = A[j][i]
		}
	}
	return B
}

func main() {

}
