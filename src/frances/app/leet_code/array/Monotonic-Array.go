package main

func isMonotonic(A []int) bool {
	var flag int
	ret := true
	for i:=1; i<len(A); i++ {
		if A[i]-A[i-1] > 0 {
			if flag == 2 {
				ret = false
				break
			}
			flag = 1
		}

		if A[i]-A[i-1] < 0 {
			if flag == 1 {
				ret = false
				break
			}
			flag = 2
		}
	}
	return ret
}

func main() {
	
}
