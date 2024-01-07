package composite

import (
	"fmt"
	"strings"
)

func DemoCollections() {
	// Array
	// primesArray := [3]int{2, 3, 5}
	primesArray := [...]int{2, 3, 5}
	fmt.Println(primesArray)
	// Slice
	// primesSlice := []int{2, 3, 5}
	primesSlice := make([]int, 3, 4)
	primesSlice = append(primesSlice, 7)
	anotherPrimesSlice := make([]int, 11, 13)
	primesSlice = append(primesSlice, anotherPrimesSlice...)
	fmt.Println(primesSlice)
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))
	for _, value := range primesSlice {
		fmt.Println(value)
	}
	// Map
	pointsMap := map[string]Point{
		"A": {X: 1, Y: 2},
		"B": {X: 1, Y: 2},
	}
	delete(pointsMap, "B")
	pointA, ok := pointsMap["A"]
	// Check that the Key does exist in the Map
	if ok {
		fmt.Println(pointA, ok)
	}
	// String
	stringSlice := []string{"B", "r", "u", "n", "o"}
	var stringBuilder strings.Builder
	for i := range stringSlice {
		// Strings are immutable, so the StringBuilder prevents each Iteration from Creating a new String
		stringBuilder.WriteString(stringSlice[i])
	}
	fmt.Printf("%v\n", stringBuilder.String())
}
