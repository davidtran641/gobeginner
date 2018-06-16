package main

import "fmt"

func add(x int, y int) int {
	return x+y
}

func main() {
	fmt.Println("8+9 = ", add(8,9))
}