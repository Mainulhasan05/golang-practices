package main

import "fmt"

// TODO: Implement these functions
// 1. Write a function 'findDuplicates' that takes a slice of integers and returns a slice of integers that appear more than once
// 2. Write a function 'wordFrequency' that takes a slice of strings and returns a map with word counts
// 3. Write a function 'mergeAndSort' that takes two sorted slices and merges them into a single sorted slice

func findDuplicates(numbers []int) []int {
	// TODO: Implement this function
	// Hint: Use a map to keep track of number occurrences
	return nil
}

func wordFrequency(words []string) map[string]int {
	// TODO: Implement this function
	// Create a map to store word counts
	return nil
}

func mergeAndSort(slice1, slice2 []int) []int {
	// TODO: Implement this function
	// Remember: The input slices are already sorted
	return nil
}

func main() {
	// Test findDuplicates
	numbers := []int{1, 2, 3, 2, 4, 5, 5, 6}
	fmt.Println("Finding duplicates in:", numbers)
	fmt.Println("Duplicates:", findDuplicates(numbers))

	// Test wordFrequency
	words := []string{"apple", "banana", "apple", "cherry", "date", "banana"}
	fmt.Println("\nCounting word frequency in:", words)
	fmt.Println("Word counts:", wordFrequency(words))

	// Test mergeAndSort
	slice1 := []int{1, 3, 5, 7}
	slice2 := []int{2, 4, 6, 8}
	fmt.Println("\nMerging and sorting:", slice1, "and", slice2)
	fmt.Println("Merged result:", mergeAndSort(slice1, slice2))
} 