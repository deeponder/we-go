package main


func largeGroupPositions(S string) [][]int {
	ret := make([][]int, 0)

	count := 1
	pos := 0
	i := 0
	for i=0; i<len(S); i++ {
		if S[i] == S[i+1] {
			count ++
		} else {
			if count >= 3 {
				tmp := []int{pos, pos+count-1}
				ret = append(ret, tmp)
			}
			pos = i+1
			count = 1
		}
	}

	if count >= 3{
		if count >= 3 {
				tmp := []int{pos, pos+count-1}
				ret = append(ret, tmp)
			}
	}

	return ret
}

func main() {
	
}
