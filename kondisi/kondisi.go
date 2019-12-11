package main

import "fmt"

func main() {
	angka := 10
	convertAngkaToString := fmt.Sprintln(angka) // Convert Int To String
	// Convert float To int int(10.0) -> Menjadi int
	fmt.Print(convertAngkaToString)

	a, b := 20, 13
	fmt.Println(a, b)
	fmt.Printf("value anda adalah %v %v %9.f \n", a, b, 22421312312.2422)

	c := 5
	if c < 10 && c > 10 || c <= 10 {
		fmt.Println("Oke")
	}

	// switch angka := 4; angka {
	// case 0:
	// 	fmt.Println("0")
	// case 4:
	// 	fmt.Println("4")
	// default:
	// 	fmt.Println("Default")
	// }
}
