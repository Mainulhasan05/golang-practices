package main

import (
	"fmt"
	"time"
)

func printPattern(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		sum += i
	}
	return sum
}

func main() {
	start := time.Now()
	fmt.Println(printPattern(100000))
	fmt.Println(time.Since(start))

}
