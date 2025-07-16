package main

import (
	"fmt"
	"sync"
	"time"
)

func process(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing job %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Finished job %d\n", id)
}

func main(){
	var wg sync.WaitGroup
	noOfTask := 10

	for i:=0; i<noOfTask; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All jobs done")
}