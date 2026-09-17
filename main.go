package main

import (
	"fmt"
	// "strings"
)

// func main() {
// 	var color string

// 	fmt.Print("masukan warna lampu lalu lintas : ")
// 	fmt.Scanln(&color)

// 	switch color {
// 	case "merah":
// 		fmt.Println("lampu merah, berhenti")
// 	case "kuning":
// 		fmt.Println("lampu kuning, bersiap-siap")
// 	case "hijau":
// 		fmt.Println("lampu hijau, jalan")
// 	default:
// 		fmt.Println("warna lampu tidak valid")
// 	}
// 		Hewan := "kucing, Ayam, Bebek, Sapi, Kambing"
// 		filter := strings.Split(Hewan, ", ")
// 		fmt.Println(filter)
// }

func main() {
	var angka int
	for {
		fmt.Print("Masukan angka : ")
		fmt.Scanln(&angka)

		if angka != int() {
			fmt.Println("yang bener kocak")
			break
		}

		if angka <= 0 {
			fmt.Println("Angka harus 1 >=")
			fmt.Print("Masukan angka lagi : ")
			fmt.Scanln(&angka)
		}

		if angka%2 == 0 {
			fmt.Println("Angka", angka, "adalah bilangan genap")
		} else if angka%2 != 0 {
			fmt.Println("Angka", angka, "adalah bilangan ganjil")
		}

	}
}
