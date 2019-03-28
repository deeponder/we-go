package main

func twoSum(numbers []int, target int) []int {
	//key number, value index
	tmpMap := make(map[int]int, len(numbers))
	for k,v := range numbers {
		if tmpMap[target-v] != 0 {
			return []int{tmpMap[target-v], k+1}
		}
		tmpMap[v] = k+1
	}
	return []int{}
}

func main() {
	
}
