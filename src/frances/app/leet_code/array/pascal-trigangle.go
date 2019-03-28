package main

//A[k][n]=A[k-1][n-1]+A[k-1][n]
func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	//arr[0][0] = 1
	for i:=0; i<numRows; i++ {
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		arr[i][i] = 1
		for j:=1; j<i; j++ {
			arr[i][j] = arr[i-1][j-1]+arr[i-1][j]
		}
	}
	return arr
}

func main() {
	
}
