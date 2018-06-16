package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["home"] = Vertex {40.1, 16.2}

	fmt.Println(m["home"])
}