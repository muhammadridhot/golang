package main

import (
	"fmt"
	"strings"
)

func main() {
	// var names [3]string
	// names[0] = "Tri"
	// names[1] = "Bima"
	// names[2] = "Ridho"

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	// for _, nama := range names {
	// 	fmt.Printf("Nama anda adalah %s \n", nama)
	// }

	// fruits := [...]string{"Apel", "Nanas", "Angur"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("Nama %d buah adalah %s \n", i+1, fruit)
	// }
	kal := "   Bang Dani"
	okeaja := " Makan"
	// fmt.Println(strings.Split(kal, ""))
	potong(kal, okeaja)
}

func potong(kal, okeaja string) {
	fmt.Println(strings.TrimSpace(kal))
	fmt.Println(strings.TrimSpace(okeaja))
}
