/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
*/
package main

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortest := strs[0]
	for _, v := range strs {
		if len(v) < len(shortest) {
			shortest = v
		}
	}

	return shortest
}

func longestCommonPrefix(strs []string) string {
	shortest := shortest(strs)
	for i, v := range shortest {
		for _, v2 := range strs {
			if v2[i] != byte(v) {
				return shortest[0:i]
			}
		}
	}

	return shortest
}

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
