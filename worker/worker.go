package worker

import (
	"strings"
	"sync"
)

// WorkerPool represents a pool of workers that can process jobs in parallel.
type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	//Done is a channel that is closed when all the workers have finished processing the jobs and the results have been collected.
	Done chan struct{}
}

// New is a constructor function that creates a new WorkerPool with the specified maximum number of workers.
func New(maxWorkers int) *WorkerPool {
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	return &WorkerPool{
		workersCount: maxWorkers,
		jobs:         make(chan Job, maxWorkers),
		results:      make(chan Result, maxWorkers),
		Done:         make(chan struct{}),
	}
}

// Run starts the worker pool and blocks until all the jobs have been processed and the results collected.
func (p *WorkerPool) Run() {
	var wg sync.WaitGroup
	wg.Add(p.workersCount)

	for i := 0; i < p.workersCount; i++ {
		// fan out worker goroutines reading from jobs channel and pushing result into results channel
		go func() {
			defer wg.Done()

			for job := range p.jobs {
				p.results <- job.execute()
			}
		}()
	}

	wg.Wait()
	close(p.Done)
	close(p.results)
}

// Results returns a read-only channel that provides access to the results produced by the worker pool.
func (p *WorkerPool) Results() <-chan Result {
	return p.results
}

// GenerateFromUrls  is a function that generates jobs from a slice of URLs and pushes them into the jobs channel
func (p *WorkerPool) GenerateFromUrls(urls []string) {
	for i, _ := range urls {
		url := urls[i]
		// Check if the URL starts with "http://" or "https://", and add it if not
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		p.jobs <- Job(url)
	}
	close(p.jobs)
}
