package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var i int

func PlayPauseRoutine() {

	var wg sync.WaitGroup

	wg.Add(1)
	command := make(chan string)
	go subroutine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}

func subroutine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var status = "play"

	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func work() {
	time.Sleep(250 * time.Microsecond)
	i++
	fmt.Println(i)
}
