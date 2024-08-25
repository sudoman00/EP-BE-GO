package main

import "fmt"

func main() {
	var s string

	fmt.Print("Enter a string: ")
	fmt.Scan(&s)

	fmt.Println(len(s)%2 == 0)
}
