package main

func numPairsDivisibleBy60(time []int) int {
	num := 0
	for i:=0; i<len(time);i++ {
		for j := i+1; j<len(time); j++ {
			if (time[i]+time[j])%60==0{
				num++
			}
		}
	}

	return num
}

func main() {
	println(numPairsDivisibleBy60([]int{30,20,150,100,40}))
}
