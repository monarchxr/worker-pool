package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- time.Duration, workerwg *sync.WaitGroup) {
	defer workerwg.Done()
	for task := range tasks {
		start := time.Now()
		fmt.Printf("Worker %d processing task %d\n", id, task)
		duration := time.Duration(rand.Intn(41)+10) * time.Millisecond
		time.Sleep(duration)
		results <- time.Since(start)
	}
}

func main() {
	startTime := time.Now()
	// so this will be the worker pool model

	// what happened in the naive concurrency model
	// was that
	// it spawns N number of goroutines at the same time
	// (not literally, there is a gap of nanoseconds)

	// so 100, 500, 1000 goroutines
	// 1 goroutine takes about 2KB size
	// if N keeps increasing
	// it can reach GBs, which is overall very poor space utilization

	// in the worker pool model
	// instead of spawning tons of goroutines at the same time
	// we spawn a certain number of workers

	// lets say 50 workers
	total_workers := 10

	// and lets say we have 500 tasks
	total_tasks := 500

	// now these workers take jobs from a job queue/ channel
	// whose size is equal to num of tasks
	tasks := make(chan int, total_tasks)

	// and a results queue, to store latency
	results := make(chan time.Duration, total_tasks)

	// create a waitgroup for workers
	var workerwg sync.WaitGroup

	// start the workers
	for i := 0; i < total_workers; i++ {
		workerwg.Add(1)
		go worker(i, tasks, results, &workerwg)
	}

	for j := 0; j < total_tasks; j++ {
		tasks <- j
	}

	close(tasks)

	go func() {
		workerwg.Wait()
		close(results)
	}()

	total_duration := 0 * time.Millisecond
	min_duration := 1000 * time.Millisecond
	max_duration := 0 * time.Millisecond

	for result := range results {
		total_duration += time.Duration(result)
		min_duration = min(min_duration, result)
		max_duration = max(max_duration, result)
	}

	fmt.Printf("\nAverage latency = %v\n", (total_duration)/time.Duration(total_tasks))
	fmt.Printf("Minimum time taken for a task = %v\n", min_duration)
	fmt.Printf("Maximum time taken for a task = %v\n", max_duration)

	endTime := time.Since(startTime)
	fmt.Printf("\nTotal wall clock time taken = %v", endTime)

}
