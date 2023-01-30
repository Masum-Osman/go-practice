package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wgroup sync.WaitGroup
var count int32

func AtomicFunc() {
	wgroup.Add(3)

	go increment("Python")
	go increment("Golang")
	go increment("JavaScript")

	wgroup.Wait()
	fmt.Println("Counter:", count)
}

func increment(name string) {
	defer wgroup.Done()

	for range name {
		atomic.AddInt32(&count, 1)
		// count++
		fmt.Println(count)
		runtime.Gosched()
	}
}

//commad:
// go run -race main.go
