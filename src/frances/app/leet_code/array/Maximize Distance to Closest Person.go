package main

func maxDistToClosest(seats []int) int {
	posHad := make([]int, 0)
	for i, v := range seats {
		if v == 1 {
			posHad = append(posHad, i)
		}
	}

	max := 0
	for i:=0; i<len(posHad)-1; i++ {
		tmp := posHad[i+1] - posHad[i]
		if tmp > max {
			max = tmp
		}
	}

	tmp2 := len(seats)
	if tmp2-posHad[len(posHad)-1] > max {
		max = tmp2-posHad[len(posHad)-1]
	}

	return max
}

func main() {
	
}
