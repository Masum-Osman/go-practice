package panic_recover

import (
	"errors"
	"fmt"
)

var result = 1

func PanicRecovery(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	// defer recover()

	if n == 0 {
		panic(errors.New("cannot multiply a number by zero"))
	} else {
		result += n
		fmt.Println("\nOutput: ", result)
	}

}
