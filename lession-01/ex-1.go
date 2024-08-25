package main

import (
	"fmt"
)

func perimeter(width, height float64) float64 {
	return (width + height) * 2
}

func area(width, height float64) float64 {
	return width * height
}

func main() {
	var width, height float64

	fmt.Print("Enter width: ")
	fmt.Scan(&width)

	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	fmt.Println("Rectangle perimeter: ", perimeter(width, height))
	fmt.Println("Rectangle area: ", area(width, height))
}
