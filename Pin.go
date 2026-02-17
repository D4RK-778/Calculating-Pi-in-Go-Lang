package main

// import
import (
	"fmt"
	"math"
)

// variables
var diameter float64 = 31.83
var perimeter float64 = 100.0

// function to calculate
func calculate(diameter float64, perimeter float64) {
	cal := perimeter / diameter
	real := math.Pi
	diference := math.Abs(real - cal)

	fmt.Println(cal, real, diference)
}

// main function
func main() {
	calculate(diameter, perimeter)
}

// Go Lang is a cool language
