package main

import "fmt"

func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

func main() {
	fmt.Println("No bottles of beer on the wall")
	receivedChan := Count(1, 99)
	// loop all until channel is close, we can use wait group
	for i := range receivedChan {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
		// Pass it around, put one up, 1 bottles of beer on the wall
		// Pass it around, put one up, 2 bottles of beer on the wall
		// ...
		// Pass it around, put one up, 99 bottles of beer on the wall
	}

	fmt.Println(100, "bottles of beer on the wall")
}
