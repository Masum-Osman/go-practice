package goroutines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func ResponseSize(url string, nums chan int) {

	defer wg.Done()

	// fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Step 2: ", url)
	defer response.Body.Close()

	// fmt.Println("Step 3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Step 4: ", len(body))
	nums <- len(body)
}

func UseGoRoutineWithWG() {

	nums := make(chan int)
	wg.Add(1)
	fmt.Println("Started Goroutines...")

	go ResponseSize("https://www.golangprograms.com", nums)
	// go ResponseSize("https://coderwall.com", nums)
	// go ResponseSize("https://stackoverflow.com", nums)
	// time.Sleep(5 * time.Second)
	fmt.Println(<-nums)

	wg.Wait()
	fmt.Println("Terminating Goroutine")
	close(nums)
}
