/*
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
 */
package main

func strStr(str, need string) int {
	nlen, strlen := len(need), len(str)
	for i:=0; i < strlen - nlen; i++ {
		if str[i:i+nlen] == need {
			return i
		}
	}
	return -1
}

func main() {
	println(strStr("hello", "la"))
}
