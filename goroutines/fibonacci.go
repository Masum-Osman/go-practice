package goroutines

import "fmt"

func FibonacciRoutine() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println("FROM CHANNEL: ", <-ch)
		}
		quit <- false
	}(n)
	// fmt.Println(<-ch, <-quit)
	fibonacci(ch, quit)
}

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
