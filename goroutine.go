package main

import (
	"log/slog"
	"time"
)

var count int

func main() {
	start := time.Now()
	defer slog.Info("time", slog.Duration("elapsed", time.Since(start)))

	ch := make(chan struct{}, 10000)
	done := make(chan struct{})

	go increment(ch, done)
	for i := 0; i < 100000; i++ {
		ch <- struct{}{}
	}
	close(ch)
	<-done

	slog.Info("Done:", "count", count)
}

func increment(ch chan struct{}, done chan struct{}) {
	for range ch {
		count++
	}
	done <- struct{}{}
}
