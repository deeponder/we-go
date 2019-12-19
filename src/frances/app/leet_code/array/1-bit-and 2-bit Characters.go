package main

// 0 | 10 11  贪婪算法
//若是1， 则一定占两位
//最后是0结尾
// https://leetcode.com/problems/1-bit-and-2-bit-characters/
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i<len(bits)-1 {
		if bits[i] == 0 {
			i++
		}else {
			i += 2
		}
	}

	if i == len(bits)-1 {
		return true
	}
	return false
}

func main() {
	
}
