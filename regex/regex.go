package regex

import (
	"fmt"
	"regexp"
)

func nonAlphaQuote() {
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern: %v", re.String())
	fmt.Println(re.MatchString(str1))

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

func RunRegexPackage() {
	nonAlphaQuote()
}
