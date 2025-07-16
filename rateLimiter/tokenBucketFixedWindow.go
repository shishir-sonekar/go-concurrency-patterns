/**
Use Case:
Allow a function to be called no more than N times per second — e.g., 2 requests/second per user or per service.
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i:=1; i<=5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Millisecond)

	for req := range requests {
		<- limiter	// throttle here
		fmt.Printf("Processing request at ", time.Now().Format("15:04:14"), "for", req, "\n")
	}
}