package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var angka int
	// fmt.Print("Masukkan angka :")
	// fmt.Scan(&angka)
	// fmt.Print("Angka yang anda masukkan adalah ", angka)
	// var angka1, angka2, hasil int

	// fmt.Print("Masukkan angka 1 : ")
	// fmt.Scan(&angka1)
	// fmt.Print("Masukkan Angka 2 : ")
	// fmt.Scan(&angka2)
	// hasil = angka1 + angka2
	// fmt.Println("Hasilnya adalah ", hasil)
	// fmt.Print("Apakah anda ingin melanjutkan ?")

	for i := 0; i < 1; i++ {
		var kal string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		kal = scanner.Text()
		fmt.Print("Kalimat yang anda masukkan yaitu ", kal)
	}
}
