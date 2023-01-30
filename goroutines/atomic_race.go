package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wgroup sync.WaitGroup
var count int32
var mutex sync.Mutex

func AtomicFunc() {
	wgroup.Add(3)

	go incrementInMutex("Python")
	go incrementInMutex("Golang")
	go incrementInMutex("JavaScript")

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

func incrementInMutex(lang string) {
	defer wgroup.Done()

	for k := 0; k < 3; k++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			count++
			fmt.Println(count)
		}
		mutex.Unlock()
	}
}

//commad:
// go run -race main.go
