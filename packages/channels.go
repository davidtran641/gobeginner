package main
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Println(s, sum)
	}
	c <- sum
}

func main() {
	s := []int {7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	fmt.Println("get result from channel")
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}