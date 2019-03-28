package main

import (
	"fmt"
	"log"
	"io/ioutil"
)

func majorityElement(nums []int) int {
	var majNum int
	count := 0
	for i:=0; i<len(nums);i++ {
		if count == 0 {
			majNum = nums[i]
			count++
		}else {
			if majNum != nums[i] {
				count --
			} else {
				count ++
			}
		}
	}
	return majNum
}

func main() {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
