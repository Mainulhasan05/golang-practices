package main

import "fmt"

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func bubbleSort(arr []int) []int {
	n := len(arr) // n = 5
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ { // j< 2
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
			printArray(arr)
		}
		if !flag {
			break // No swaps means the array is sorted
		}
		// printArray(arr) // Print the array after each pass
	}
	return arr
}

func main() {
	arr := []int{31, 12, 17, 5, 2}
	fmt.Println("Before sorting:", arr)
	sortedArr := bubbleSort(arr)
	fmt.Println("After sorting:", sortedArr)
}
