package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var id int
	var amount float32
	var name string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the id: ")
	scanner.Scan()
	id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Enter the name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter the amount: ")
	scanner.Scan()
	amountFloat, _ := strconv.ParseFloat(scanner.Text(), 32)
	amount = float32(amountFloat)

	fmt.Println("Id:", id)
	fmt.Println("Name:", name)
	fmt.Println("Amount:", amount)
}
