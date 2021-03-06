package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

// Hello to return hello string
func Hello(name string) string {
	return "Hello, " + name
}
