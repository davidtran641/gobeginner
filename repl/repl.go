package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	loop()
}

func loop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		s := string(bytes)

		result, err := parse(s)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Println(result)

	}
}

func parse(s string) (float64, error) {
	arr := strings.Split(strings.Trim(s, " "), " ")
	if len(arr) != 3 {
		return 0, errors.New("Invalid expression")
	}

	x, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return 0, err
	}

	y, err := strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return 0, err
	}

	switch arr[1] {
	case "+":
		return AddOperator{}.calculate(x, y)
	case "-":
		return SubOperator{}.calculate(x, y)
	case "*":
		return MulOperator{}.calculate(x, y)
	case "/":
		return DivOperator{}.calculate(x, y)
	}
	return 0, errors.New("Invalid operation")
}

type Operator interface {
	calculate(x float64, y float64) (float64, error)
}

type AddOperator struct{}

func (op AddOperator) calculate(x float64, y float64) (float64, error) {
	return x + y, nil
}

type SubOperator struct{}

func (op SubOperator) calculate(x float64, y float64) (float64, error) {
	return x - y, nil
}

type MulOperator struct{}

func (op MulOperator) calculate(x float64, y float64) (float64, error) {
	return x * y, nil
}

type DivOperator struct{}

func (op DivOperator) calculate(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return x / y, nil
}
