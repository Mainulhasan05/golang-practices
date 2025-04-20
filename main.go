package main

import "fmt"

func test() {
	fmt.Println("Test function")
	fmt.Println("Sum:", add(30, 4))
}

func main() {
	fmt.Println("Hello, World!")
	test()
	fmt.Println("Sum:", add(3, 4))
}
func add(a, b int) int {

	return a + b
}
