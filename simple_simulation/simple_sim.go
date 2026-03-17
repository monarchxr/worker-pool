package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processTask() {
	// here we need to process task

	// first we'll sleep the task for a random amount of time
	// we do this because
	// in reality tasks dont just compute, they wait
	// wait for - disk read,writes, network respones, api calls, db queries

	// waiting time is where concurrency actually occurs
	// while 1 goroutine is blocked, another can do work

	// lets first generate the duration
	duration := time.Duration(rand.Intn(41)+10) * time.Millisecond
	// its a duration of 10-50ms

	// now make it sleep
	time.Sleep(duration)
}

func main() {
	startTime := time.Now()

	// so this will be a sequential processor
	// only to assess how slow it will be compared to other upcoming models alright

	// we pass it to the processing function
	for i := 0; i < 500; i++ {
		processTask()
	}

	//end

	endTime := time.Since(startTime)
	fmt.Printf("End time = %v\n", endTime)
}
