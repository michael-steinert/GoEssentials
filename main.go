package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/michael-steinert/GoEssentials/calculator"
	"github.com/michael-steinert/GoEssentials/composite"
)

func main() {
	fmt.Println("Go Essentials", rand.Intn(10))
	fmt.Printf("21 + 42 = %d\n", calculator.Addition(21, 42))
	quotient, remainder, err := calculator.Divide(17, 3)
	if err == nil {
		fmt.Printf("17 / 3 = %d R %d\n", quotient, remainder)
	}
	fmt.Printf("Sum from 1 to 10 is %d\n", calculator.Sum(1, 10))
	fmt.Printf("Sum until 10 is %d\n", calculator.SumUntil(10))
	fmt.Printf("Is 16 a Square Number? %t\n", calculator.IsSquareNumber(16))
	squareRootResult, err := calculator.SquareRoot(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Square Root of 16 is %f\n", squareRootResult)
	}
	fmt.Printf("Sum from 1 to 10 is %d\n", calculator.SumFromLeftToRight(1, 10))
	fmt.Printf("Product from 1 to 10 is %d\n", calculator.MultiplyFromLeftToRight(1, 10))
	// Channel for asynchronous Code
	queue := make(chan int)
	// GoRoutine i.e. asynchronous Code
	go func(q chan int) {
		time.Sleep(time.Second * 2)
		fmt.Println("After two Seconds")
		q <- 1001
	}(queue)
	fmt.Printf("Value from Queue is %d\n", <-queue)
	fmt.Println("After two Seconds")
	// Struct
	point := composite.Point{X: 1, Y: 1}
	fmt.Printf("Point Coordinates are %d and %d\n", point.X, point.Y)
	fmt.Printf("Distance from Zero with Coordinates X:1 and Y:1 is %f\n", point.DistanceFromZero())
	print(point.DistanceFromZero())
	print(point.String())

}

func print(value interface{}) {
	// Type Assertion
	switch value.(type) {
	case float64:
		fmt.Println(value)
	case string:
		fmt.Println(value)
	default:
		fmt.Println("Value can not be printed")
	}
}
