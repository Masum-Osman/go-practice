package goroutines

import (
	"fmt"
	"strconv"
)

func RunMultipleRoutines() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Go Routine : " + strconv.Itoa(i)
			}
		}(i)
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}
}

func RoutineSenderReceiver() {
	c := generator()
	receiver(c)
}

func generator() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func receiver(c <-chan int) {
	fmt.Println(c)
	for v := range c {
		fmt.Println(v)
	}
}
