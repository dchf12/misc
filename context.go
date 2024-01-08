package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx := context.Background()
	err := runJobs(ctx)
	fmt.Println(err)
}

func runJobs(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ec := make(chan error)
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			cmd := exec.CommandContext(ctx, "sleep", "30")
			err := cmd.Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}
	go func(){
		time.Sleep(10 * time.Second)
		ec <- errors.New("application error")
	}()

	for i := 0; i < 11; i++ {
		select {
		case err := <-ec:
			return err
		case <-done:
		}
	}
	return nil
}
