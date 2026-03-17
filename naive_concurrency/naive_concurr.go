package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processTask(wg *sync.WaitGroup) {

	defer wg.Done()

	duration := time.Duration(rand.Intn(41)+10) * time.Millisecond

	time.Sleep(duration)
}

func main() {

	startTime := time.Now()

	//ok so this a naive version of concurrent processing

	// we'll be working with goroutines
	// whats a goroutine?

	// a lightweight independent thread
	// enables easy concurrency

	// how to make goroutine?
	// just prefix the function call with go

	// also we need a waitgroup before spawning each goroutine

	var wg sync.WaitGroup
	// a waitgroup is like a counter
	// formally, a semaphore

	// it waits for a group of goroutines to finish executing
	// add increments the counter by input value
	// done decrements the counter by 1
	// wait - blocks main thread until counter hits 0

	//lets begin

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go processTask(&wg)
	}
	wg.Wait()

	endTime := time.Since(startTime)
	fmt.Printf("End time = %v\n", endTime)
}
