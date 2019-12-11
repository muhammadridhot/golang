package main

import "fmt"

func main() {
	a := 10
	b := 10 + 5

	fmt.Println(a, b)

	condLeft := true
	condRight := false
	condResult := !condLeft || condRight

	fmt.Println("condLeft :", condLeft)
	fmt.Println("condRight :", condRight)
	fmt.Println("condResult :", condResult)

}
