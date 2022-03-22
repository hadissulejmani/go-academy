package main

import "fmt"

type Action func() error

func SafeExec(a Action) Action {
	return func() error {
		return fmt.Errorf("safe exec: %w")
	}
}

func main() {
	fmt.Println("yayak")
}
