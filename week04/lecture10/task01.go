package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	/* Feel free to add your own state here */
	counter     int
	lastElement string
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElement != "foo" {
				fmt.Print("foo")
				cp.lastElement = "foo"
				cp.counter++
			}
			cp.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElement != "bar" {
				fmt.Print("bar")
				cp.lastElement = "bar"
				cp.counter++
			}
			cp.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

/*foobarfoobarfoobarfoobarfoobar*/
func main() {
	times := 10

	cp := &ConcurrentPrinter{}
	cp.lastElement = "bar"

	cp.PrintFoo(times)
	cp.PrintBar(times)

	cp.Wait()
}
