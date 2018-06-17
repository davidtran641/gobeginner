package main
import "fmt"

func fibonacci() func() int {
	s1 := 0
	s2 := 1

	return func() int {
		a1 := s1
		a := s2
		s2 = s1 + s2
		s1 = a
		return a1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}