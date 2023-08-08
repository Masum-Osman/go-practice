package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func cpuIntensiveTask() {
	for i := 0; i < 100000000; i++ {
		_ = i * i
	}
}

func pprofTest() {
	// Start the pprof HTTP server on port 8080.
	go func() {
		fmt.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// Simulate a CPU-intensive task.
	for i := 0; i < 5; i++ {
		cpuIntensiveTask()
		time.Sleep(1 * time.Second)
	}
}
