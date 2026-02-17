package main

import (
	"fmt"
	"math"
)

var diameter float64 = 31.83
var perimeter float64 = 100

func main() {
	cal := perimeter / diameter
	real := math.Pi
	diference := math.Abs(real - cal)

	fmt.Println(cal)
	fmt.Println(real)
	fmt.Println(diference)
}
