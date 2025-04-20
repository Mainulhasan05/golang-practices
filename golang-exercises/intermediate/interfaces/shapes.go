package main

import (
	"fmt"
	"math"
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Implement these structs and their methods
// 1. Create a Rectangle struct with width and height fields
// 2. Create a Circle struct with radius field
// 3. Implement the Shape interface for both structs

type Rectangle struct {
	// TODO: Add fields
}

func (r Rectangle) Area() float64 {
	// TODO: Implement area calculation
	return 0
}

func (r Rectangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 0
}

type Circle struct {
	// TODO: Add fields
}

func (c Circle) Area() float64 {
	// TODO: Implement area calculation
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 0
}

// TODO: Implement a function that takes a slice of Shapes and returns the total area
func totalArea(shapes []Shape) float64 {
	// TODO: Calculate total area of all shapes
	return 0
}

func main() {
	// TODO: Create instances of Rectangle and Circle
	// Calculate and print their areas and perimeters
	// Test the totalArea function with a slice of different shapes
	
	fmt.Println("Implement the shapes exercise!")
	
	// Example usage (uncomment and implement):
	/*
	rect := Rectangle{width: 5, height: 3}
	circle := Circle{radius: 2}
	
	shapes := []Shape{rect, circle}
	total := totalArea(shapes)
	
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())
	fmt.Printf("Total area of all shapes: %.2f\n", total)
	*/
} 