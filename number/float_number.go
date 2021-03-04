package main

import (
	"fmt"
)

func main() {
	var num float32

	num = 9876543.01
	s := fmt.Sprintf("%.2f", num) // 9876543.00
	fmt.Println(s)

	num = 987654321.01
	s = fmt.Sprintf("%.2f", num) // 987654336.00
	fmt.Println(s)
}
