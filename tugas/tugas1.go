package main

import (
	"fmt"
	"math"
)

func main() {

	jari, hasilLuasLingkaran, hasilKelilingLingkaran := 0.0, 0.0, 0.0

	sisi, hasilLuasPersegi, hasilKelilingPersegi := 0.0, 0.0, 0.0

	panjang, lebar, tinggi, hasilLuasPermukaanBalok, hasilVolumeBalok := 0.0, 0.0, 0.0, 0.0, 0.0

	var lanjutkan string

	angka := 0
	for i := 0; i < 1; i++ {
		fmt.Println("1. Luas Lingkaran")
		fmt.Println("2. Keliling Lingkaran")
		fmt.Println("3. Luas Persegi")
		fmt.Println("4. Keliling Persegi")
		fmt.Println("5. Luas Permukaan Balok")
		fmt.Println("6. Volume Balok")
		fmt.Println("Pilihlah angka sesuai keterangan diatas")
		fmt.Print("Masukkan angka : ")
		fmt.Scan(&angka)
		switch angka {
		case 1:
			//Luas Lingkaran
			//L = π × d²/4 = π × r²
			{
				fmt.Print("Masukkan Jari-Jari : ")
				fmt.Scan(&jari)
				hasilLuasLingkaran = math.Pi * (math.Pow(jari, 2))
				fmt.Printf("Luas lingkaran adalah %.2f \n", hasilLuasLingkaran)
			}

		case 2:
			//Keliling Lingkaran
			//K = π × d = 2 × π × r
			{
				fmt.Print("Masukkan Jari-Jari : ")
				fmt.Scan(&jari)
				hasilKelilingLingkaran = 2 * math.Pi * (math.Pow(jari, 2))
				fmt.Printf("Luas lingkaran adalah %.2f \n", hasilKelilingLingkaran)
			}
		case 3:
			// //Luas Persegi
			// //L = s × s
			{
				fmt.Print("Masukkan Sisi : ")
				fmt.Scan(&sisi)
				hasilLuasPersegi = math.Pow(sisi, 2)
				fmt.Printf("Luas persegi adalah %.2f \n", hasilLuasPersegi)
			}
		case 4:
			// //Keliling Persegi
			// //Kll = 4 × s
			{
				fmt.Print("Masukkan Sisi : ")
				fmt.Scan(&sisi)
				hasilKelilingPersegi = 4 * sisi
				fmt.Printf("Keliling persegi adalah %.2f \n", hasilKelilingPersegi)
			}
		case 5:
			// //Luas Permukaan Balok
			// //	L = 2 x (pl + pt + lt)
			{
				fmt.Print("Masukkan Panjang: ")
				fmt.Scan(&panjang)
				fmt.Print("Masukkan Lebar: ")
				fmt.Scan(&lebar)
				fmt.Print("Masukkan Tinggi: ")
				fmt.Scan(&tinggi)
				hasilLuasPermukaanBalok = 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
				fmt.Printf("Luar Permukaan Balok adalah %.2f \n", hasilLuasPermukaanBalok)
			}
		case 6:
			// //Volume Balok
			// //	V = p x l x t
			{
				fmt.Print("Masukkan Panjang: ")
				fmt.Scan(&panjang)
				fmt.Print("Masukkan Lebar: ")
				fmt.Scan(&lebar)
				fmt.Print("Masukkan Tinggi: ")
				fmt.Scan(&tinggi)
				hasilVolumeBalok = panjang * lebar * tinggi
				fmt.Printf("Volume Balok adalah %.2f \n", hasilVolumeBalok)
			}
		default:
			fmt.Println("Angka anda masukkan tidak sesuai keterangan, masukkan angka antara 1 sampai 6")
			// fmt.Print("Anda ingin melanjutkan (Y/T) : ")
			// fmt.Scan(&lanjutkan)
			// if lanjutkan1 == strings.ToLower("Y") {
			// 	i--
			// }
			// var kal string
			// scanner := bufio.NewScanner(os.Stdin)
			// scanner.Scan()
			// kal = scanner.Text()

			// if kal == "Y" {
			// 	i--
			// }

		}
		fmt.Print("Anda ingin melanjutkan (Y/T) : ")
		fmt.Scan(&lanjutkan)
		if lanjutkan == "Y" {
			i--
		}
	}
}
