package main

import (
	"io"
	"fmt"
	"bytes"
)

func containsDuplicate(nums []int) bool {
	cont := make(map[int]int, 0)
	for i:=0; i<len(nums); i++ {
		if cont[nums[i]] == 1 {
			return true
		} else {
			cont[nums[i]] = 1
		}
	}
	return false
}


func main() {


}

