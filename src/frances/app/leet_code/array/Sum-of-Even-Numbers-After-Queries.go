package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	//先把偶数加起来
	for _, v := range A {
		if v % 2 == 0 {
			sum += v
		}
	}
	//根据queries的情况，调整, 偶=偶+偶/奇+奇
	ret := make([]int, len(queries))
	for i:=0; i< len(queries); i++ {
		index := queries[i][1]
		val := queries[i][0]
		if val % 2 == 0 && A[index] % 2 == 0{
			sum = sum + val
		} else if val % 2 != 0 && A[index] % 2 == 0 {
			sum = sum - A[index]
		} else if val % 2 != 0 && A[index] % 2 != 0 {
			sum = sum + A[index] +val
		}

		A[index] = A[index] + val
		ret[i] =  sum
	}

	return ret
}

func main() {
	
}
