package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	/* Add other fields you might need */
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	/*Implement the rest */
	return &BufferedContext{Context: ctx, CancelFunc: cancel}
}
func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */
	/* make it so that the result channel gets closed in one of the to cases;
	   a) the emebdded context times out
	   b) the buffer gets filled
	*/
	return bc.Context.Done()
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */
	/* Implement the rest */
}

func main() {
	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 200) // try different values here
				fmt.Println("bar")
			}
		}
	})
}
