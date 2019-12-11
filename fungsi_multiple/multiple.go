package main

import (
	"fmt"
	"math"
)

func main() {
	diameter := 15.0
	area, circumference := calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)

	circumference := math.Pi * d

	return area, circumference
}
