package main

func uniqueMorseRepresentations(words []string) int {
	M := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	morseMap := make(map[string]int, 0)

	for _, word := range words {
		morse := ""
		for _, char := range word {
			morse += M[char-'a']
		}
		morseMap[morse]++
	}

	return len(morseMap)
}

func main() {
	
}
