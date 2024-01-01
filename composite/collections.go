package composite

import "fmt"

func DemoCollections() {
	// Array
	primesArray := [3]int{2, 3, 5}
	fmt.Println(primesArray)
	// Slice
	// primesSlice := []int{2, 3, 5}
	primesSlice := make([]int, 3, 4)
	primesSlice = append(primesSlice, 7)
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
	if ok {
		fmt.Println(pointA, ok)
	}
}
