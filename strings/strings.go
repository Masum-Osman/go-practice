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

func Interpret(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)

	return command
}

func InterpretLoop(command string) string {
	ans := ""

	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			ans = ans + "G"
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				ans = ans + "o"
			} else if command[i+1] == 'a' {
				ans = ans + "al"
			}
		}
	}

	return ans
}

var STACK []string

func stackLength() int {
	return len(STACK)
}

func stackPush(val string) {
	STACK = append(STACK, val)
}

func stackPop() {
	STACK = STACK[:len(STACK)-1]
}

func IsValidParentheses(s string) bool {

	for _, i := range s {
		if i == '(' || i == '{' {
			stackPush(string(i))
		} else if i == ')' || i == '}' {
			stackPop()
		}
	}

	if stackLength() == 0 {
		return true
	} else {
		return false
	}
}
