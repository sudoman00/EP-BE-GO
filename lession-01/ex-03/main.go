package main

import (
	"fmt"
	"sort"
)

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func findMinMax(numbers []float64) (float64, float64) {
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}

func findMin(numbers []float64) float64 {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func calculateAverage(sum float64, count int) float64 {
	return sum / float64(count)
}

func sortNumbers(numbers []float64) []float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	return sortedNumbers
}

func main() {
	var n int
	fmt.Print("Enter number of slice: ")
	fmt.Scan(&n)

	numbers := make([]float64, n)
	fmt.Print("Enter slide values: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	sum := calculateSum(numbers)

	min, max := findMinMax(numbers)

	average := calculateAverage(sum, len(numbers))

	sortedNumbers := sortNumbers(numbers)

	fmt.Println("Sum:", sum)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
	fmt.Println("Average:", average)
	fmt.Println("Sorted:", sortedNumbers)
}
