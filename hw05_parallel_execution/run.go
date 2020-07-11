package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(workQueue <-chan Task, syncNum *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if task, ok := <-workQueue; ok && atomic.LoadInt64(syncNum) >= 0 {
			if task() != nil {
				atomic.AddInt64(syncNum, -1)
			}
		} else {
			return
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
//
func Run(tasks []Task, N int, M int) error { //nolint
	// Initialize channels and workers
	workQueue := make(chan Task, len(tasks))
	wg := sync.WaitGroup{}
	var syncNum int64 = int64(M)
	// Start workers
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(workQueue, &syncNum, &wg)
	}
	// Fill work queue
	for _, task := range tasks {
		workQueue <- task
	}
	close(workQueue)
	wg.Wait()
	if atomic.LoadInt64(&syncNum) >= 0 {
		return nil
	}
	return ErrErrorsLimitExceeded
}
