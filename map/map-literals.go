package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"home" : {12.1, 6.4},
	"work" : {10.1, 482.1},
}

func main() {
	fmt.Println(m)
}