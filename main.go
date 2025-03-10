package main

import "fmt"

func main() {
	var id int
	var amount float32
	var name string

	fmt.Print("Enter the id:")
	fmt.Scan(&id)
	fmt.Print("Enter the name:")
	fmt.Scan(&name)
	fmt.Print("Enter the amount:")
	fmt.Scan(&amount)

	fmt.Println("Id:", id)
	fmt.Println("Name:", name)
	fmt.Println("Amount:", amount)

}
