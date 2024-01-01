package composite

import (
	"fmt"
	"math"
)

// Interface is implicit applied, when Struct meet its Requirements
type HasDistance interface {
	DistanceFromZero() float64
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Method with Value Receiver (Parameter)
func (point Point) DistanceFromZero() float64 {
	return math.Sqrt(math.Pow(float64(point.X), 2) + math.Pow(float64(point.Y), 2))
}

// Stringer Interface
func (point Point) String() string {
	return fmt.Sprintf("(%d | %d)", point.X, point.Y)
}
