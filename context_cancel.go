package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

func taskFirst(ctx context.Context) error {
	fmt.Println("start first task")
	time.Sleep(4 * time.Second)
	return errors.New("first task end with error")
}

func taskSecond(ctx context.Context) {
	fmt.Println("start second task")
	ticker := time.NewTicker(500 * time.Millisecond)
	i := 0
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("second task ticked", i)
		case <-ctx.Done():
			fmt.Println("second task done")
			return
		}
	}
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := taskFirst(ctx)

		if err != nil {
			fmt.Println(err)
			cancel()
		}
	}()

	taskSecond(ctx)
	wg.Wait()
}
