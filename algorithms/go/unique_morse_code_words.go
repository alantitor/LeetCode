package main

import (
	"fmt"
	"strings"
)

func main() {
	testCase := [...]string{"gin", "zen", "gig", "msg"}

	ans := uniqueMorseRepresentations(testCase[:])

	fmt.Println("ans: ", ans)
}

func uniqueMorseRepresentations(words []string) int {
	morseCode := [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseCodeTable := make(map[string]string)
	result := make(map[string]bool)

	for i := 0; i < 26; i++ {
		c := string('a' + i)
		morseCodeTable[c] = morseCode[i]
	}

	for _, word := range words {
		var b strings.Builder

		for _, c := range word {
			b.WriteString(morseCodeTable[string(c)])
		}
		result[b.String()] = true
	}

	return len(result)
}
