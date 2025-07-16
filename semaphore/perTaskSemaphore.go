/**
Control how many goroutines can run concurrently — even when launching many tasks — using a channel as a semaphore.
This is perfect when:
You have lots of tasks
You want to avoid overwhelming system resources
But you still want per-task goroutines for simplicity
**/

package main

import (
	"fmt"
	"time"
	"sync"
)

func process(id int, wg sync.WaitGroup, sem chan struct{}){
	defer wg.Done()

	sem <- struct{}{}

	fmt.Printf("Task %d started", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d finished", id)

	<- sem
}

func main(){
	var wg sync.WaitGroup

	maxConcurrency := 3
	noOfTask := 20

	var sem make(chan struct{}, maxConcurrency)

	for i:=0; i<noOfTask; i++ {
		wg.Add(1)
		go process(i, &wg, sem)
	}

	wg.Wait()
	fmt.Printf("Finished processing all tasks")
}