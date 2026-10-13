package main

import (
	"context"
	"sync"
	"testing"
)

func TestWorkers(t *testing.T) {
	ctx := context.Background()

	jobs := make(chan int, 5)
	results := make(chan Result, 5)

	var wg sync.WaitGroup

	// Запускаю три воркера
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Workers(ctx, jobs, results, &wg)
	}

	// Отправляю пять заданий
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Закрываю results после завершения воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	got := make(map[int]bool)

	for result := range results {
		if result.Err != nil {
			t.Fatalf("unexpected error: %v", result.Err)
		}

		got[result.Value] = true
	}

	// Проверяю, что обработаны все пять заданий
	expected := []int{2, 4, 6, 8, 10}

	for _, value := range expected {
		if !got[value] {
			t.Errorf("missing result: %d", value)
		}
	}

	if len(got) != len(expected) {
		t.Errorf("got %d unique results, want %d", len(got), len(expected))
	}
}
