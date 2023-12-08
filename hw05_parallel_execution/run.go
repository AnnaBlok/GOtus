package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 0 {
		m = len(tasks) + 1
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for ; n > 0; n-- {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if m == 0 || len(tasks) == 0 {
					mu.Unlock()
					break
				}
				task := tasks[0]
				tasks = tasks[1:]
				mu.Unlock()
				err := task()
				mu.Lock()
				if err != nil && m > 0 {
					m--
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	if m == 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
