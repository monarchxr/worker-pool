package main

import (
	"math/rand"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func worker(id int, tasks <-chan int, results chan<- time.Duration, workerwg *sync.WaitGroup) {
	defer workerwg.Done()
	for range tasks {
		start := time.Now()
		// fmt.Printf("Worker %d processing task %d\n", id, task)
		duration := time.Duration(rand.Intn(41)+10) * time.Millisecond
		time.Sleep(duration)
		results <- time.Since(start)
	}
}

func runWorkerPool(total_workers int) time.Duration {
	startTime := time.Now()
	total_tasks := 500

	tasks := make(chan int, total_tasks)

	results := make(chan time.Duration, total_tasks)

	var workerwg sync.WaitGroup

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

	for range results {
	}

	endTime := time.Since(startTime)
	return endTime

}

func plotResults(workerSizes [11]int, durations [11]float64) {
	pts := make(plotter.XYs, 11)

	for i := range pts {
		pts[i].X = float64(workerSizes[i])
		pts[i].Y = durations[i]
	}

	p := plot.New()

	p.Title.Text = "Workers vs Time taken"
	p.X.Label.Text = "Number of workers"
	p.Y.Label.Text = "Time taken in milliseconds"

	plotutil.AddLinePoints(p, "", pts)
	p.Save(6*vg.Inch, 4*vg.Inch, "results.png")
}

func main() {

	workerSizes := [11]int{1, 10, 20, 25, 50, 75, 100, 250, 300, 400, 500}
	durations := [11]float64{}

	for i, w := range workerSizes {
		d := runWorkerPool(w)
		durations[i] = float64(d.Milliseconds())
	}

	plotResults(workerSizes, durations)

}
