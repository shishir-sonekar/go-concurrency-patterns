package main 

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID		int
	data	int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing started for job %d\n", id, job.ID)
		time.Sleep(time.Second*1)
		fmt.Printf("Worker %d processing finished for job %d\n", id, job.ID)
	}
}

func main() {
	jobCount := 5
	workerCount := 3

	jobs := make(chan Job)

	var wg sync.WaitGroup
	
	// create workers
	for w:=0; w<workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// create jobs
	for j:=0; j<jobCount; j++ {
		jobs <- Job{ ID: j, data: j*10 }
	}

	close(jobs)
	wg.Wait()
	fmt.Println("All jobs finished")
}