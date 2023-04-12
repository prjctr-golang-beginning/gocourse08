package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case val := <-ch1:
				fmt.Println("Received from ch1:", val)
				ch2 <- val
			}
		}
	}()

	go func() {
		for {
			select {
			case val := <-ch2:
				fmt.Println("Received from ch2:", val)
				ch1 <- val
			}
		}
	}()

	ch1 <- 1 // sending a value to channel ch1

	// program will deadlock here, as both goroutines are waiting for values from each other
}
