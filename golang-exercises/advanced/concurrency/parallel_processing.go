package main

import (
	"fmt"
	"time"
)

// TODO: Implement these functions
// 1. Create a function that generates numbers and sends them to a channel
// 2. Create a function that receives numbers and calculates their squares
// 3. Implement a worker pool pattern for processing numbers in parallel

// NumberGenerator sends numbers to the channel
func generateNumbers(count int, ch chan<- int) {
	// TODO: Generate 'count' numbers and send them to the channel
	// Remember to close the channel when done
}

// SquareCalculator receives numbers and returns their squares
func calculateSquares(input <-chan int, output chan<- int) {
	// TODO: Receive numbers from input channel, calculate squares,
	// and send results to output channel
}

// WorkerPool creates a pool of worker goroutines
func createWorkerPool(numWorkers int, input <-chan int, output chan<- int) {
	// TODO: Create numWorkers goroutines that process numbers
	// Each worker should run the calculateSquares function
}

func main() {
	// TODO: Set up channels and implement the following workflow:
	// 1. Generate numbers using generateNumbers
	// 2. Process them in parallel using a worker pool
	// 3. Collect and print results
	
	fmt.Println("Implement the parallel processing exercise!")
	
	// Example usage (uncomment and implement):
	/*
	inputChan := make(chan int, 100)
	outputChan := make(chan int, 100)
	
	// Generate numbers in a separate goroutine
	go generateNumbers(10, inputChan)
	
	// Create worker pool
	createWorkerPool(3, inputChan, outputChan)
	
	// Collect and print results
	// Remember to handle channel closing and completion
	*/
}

// Optional Bonus Challenges:
// 1. Add error handling using additional channels
// 2. Implement a timeout mechanism
// 3. Add a context for cancellation
// 4. Implement rate limiting 