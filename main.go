package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Workers(ctx context.Context, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-jobs:
			if !ok {
				return
			}

			timer := time.NewTimer(300 * time.Millisecond)

			select {
			case <-timer.C:
			case <-ctx.Done():
				timer.Stop()
				return
			}

			res := Process(v)
			select {
			case results <- res:
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan Result, 3)
	workers := 3
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	for y := 1; y <= workers; y++ {
		wg.Add(1)
		go Workers(ctx, jobs, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i:
			case <-ctx.Done():
				return
			}
		}

	}()

	var sum int
	var count int
	for read := range results {
		if read.Err != nil {
			fmt.Printf("Ошибка обработки:%d,%s\n", read.Value, read.Err)
			count++
			continue
		}
		sum += read.Value
		fmt.Printf("Число:%d\n", read.Value)
	}
	fmt.Printf("Error:%d\n", count)
	fmt.Printf("SUM:%d\n", sum)
}
