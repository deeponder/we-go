package main

import "strings"

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	for k, v := range sArr {
		sArr[k] = reverseStr(v)
	}

	return strings.Join(sArr, " ")
}

func reverseStr(str string) string {
	runes := []rune(str)
	for i, j:=0, len(runes)-1; i<j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	
}
