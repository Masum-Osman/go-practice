package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var gr sync.WaitGroup

func RunChannels() {
	rand.Seed(time.Now().Unix())

	projects := make(chan string, 10)

	gr.Add(5)

	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}
	fmt.Println(projects)

	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	close(projects)
	gr.Wait()
}

func employee(projects chan string, employee int) {
	fmt.Println("1st line: ---", projects, employee)
	defer gr.Done()

	for {
		project, result := <-projects

		if !result {
			fmt.Printf("Employee : %d :Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started %s\n", employee, project)

		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("\nTime to sleep", sleep, "ms")

		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}
}
