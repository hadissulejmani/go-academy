package main

import (
	"fmt"
	"sync"
)

func processEven(inputs []int) chan int {
	response := make(chan int, len(inputs))
	var wg sync.WaitGroup

	for _, v := range inputs {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				response <- i
			}
		}(v)
	}
	wg.Wait()
	close(response)
	return response
}

func processOdd(inputs []int) chan int {
	responseChannel := make(chan int, len(inputs))
	var w sync.WaitGroup

	for _, v := range inputs {
		w.Add(1)
		go func(j int) {
			defer w.Done()
			if j%2 != 0 {
				responseChannel <- j
			}
		}(v)
	}
	w.Wait()
	close(responseChannel)
	return responseChannel
}

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}

	fmt.Println("Even numbers:")
	evenCh := processEven(inputs)
	// range over the channel and print the results
	for v := range evenCh {
		fmt.Println(v)
	}

	fmt.Println("Odd numbers:")
	oddCh := processOdd(inputs)
	// range over the channel and print the results
	for v := range oddCh {
		fmt.Println(v)
	}
}
