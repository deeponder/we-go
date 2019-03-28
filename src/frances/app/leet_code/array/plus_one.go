package main

func plusOne(digits []int) []int {
	lenth := len(digits)
	if lenth == 0 {
		return []int{1}
	}

	digits[lenth-1]++

	//进位
	for i:=lenth-1; i>0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1] += 1
	}

	//首位进位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main()  {
	plusOne([]int{7,2,8,5,0,9,1,2,9,5,3,6,6,7})
}