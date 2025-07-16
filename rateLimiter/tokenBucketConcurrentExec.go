package main

import (
	"fmt"
	"time"
	"sync"
)

func limitWorker(id int, wg *sync.WaitGroup, limiter chan time.Time) {
	defer wg.Done()

	<-limiter
	fmt.Println("Processing task", id, "at", time.Now().Format("12:00:00:000"))
}

func main(){
	var wg sync.WaitGroup
	limiter := make(chan time.Time, 2)

	go func(){
		ticker := time.NewTicker(500*time.Millisecond)
		defer ticker.Stop()
		for t := range ticker.C {
			select{
			case limiter <- t:
			default:	// limiter full - ignore
			}
		}
	}()

	for i:=1; i<=20; i++ {
		wg.Add(1)
		go limitWorker(i, &wg, limiter)
	}

	wg.Wait()
	fmt.Println("Finished processing tasks")
}