package asynchronous

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func WaitGroup() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 5; i++ {
		// Increasing WaitGroup outside asynchronous Function to avoid Race Condition
		waitGroup.Add(1)
		go func(i int) {
			count(fmt.Sprintf("Async Function Call %d", i))
			// Decreasing WaitGroup
			waitGroup.Done()
		}(i)
	}
	// Waiting until all asynchronous Function in WaitGroup are finished
	waitGroup.Wait()
}

func count(prefix string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("%s: %s\n", prefix, strconv.Itoa(i))
		time.Sleep(time.Millisecond * 500)
	}
}
