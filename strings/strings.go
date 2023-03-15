package strings

import (
	"strings"
	"unicode"
)

func TrimString() string {
	str := "\t Hello, World\n "
	// fmt.Printf("Before Trim Length: %d String:%v\n", len(str), str)
	trim := strings.TrimSpace(str)
	// fmt.Printf("After Trim Length: %d String:%v\n", len(trim), trim)

	return trim
}

func IsOnlyLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func AllUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func IfContains(w, l string) bool {
	return strings.Contains(w, l)
}

func DefangIPaddr(addr string) string {
	return strings.Replace(addr, ".", "[.]", -1)
}

func FinalValueAfterOperations(operations []string) int {
	x := 0

	for _, i := range operations {
		if i == "++X" || i == "X++" {
			x++
		} else if i == "--X" || i == "X--" {
			x--
		}
	}
	return x
}

func NumJewelsInStones(jewels string, stones string) int {

	count := 0

	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if stones[i] == jewels[j] {
				count++
			}
		}
	}

	return count
}

func MostWordsFound(sentences []string) int {
	wordCounter := 0

	for _, r := range sentences {
		splitted := strings.Split(r, " ")
		if len(splitted) > wordCounter {
			wordCounter = len(splitted)
		}
	}
	return wordCounter
}
