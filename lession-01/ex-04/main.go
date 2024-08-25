package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		remaining := target - num
		if index, found := hashMap[remaining]; found {
			return []int{index, i}
		}

		hashMap[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println("Result:", result)
}
