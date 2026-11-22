package main

import (
	"fmt"
	"math"
)

func Panjang(p, l float64, luas, keliling *float64) {
	*luas = p * l
	*keliling = 2 * (p + l)
}

const phi = 3.14

func Lingkaran(r float64, luas, keliling *float64) {
	*luas = phi * r * r
	*keliling = 2 * phi * r
}

func Segi3(a, t float64, luas, keliling *float64) {
	miring := math.Sqrt(a*a + t*t)
	*luas = 0.5 * a * t
	*keliling = a + t + miring
}

func main() {
	var jenis string
	var a, b int
	var luas, keliling float64

	fmt.Scan(&jenis)

	switch jenis {
	case "P":
		fmt.Scan(&a, &b)
		Panjang(float64(a), float64(b), &luas, &keliling)
	case "L":
		fmt.Scan(&a)
		Lingkaran(float64(a), &luas, &keliling)
	case "S":
		fmt.Scan(&a, &b)
		Segi3(float64(a), float64(b), &luas, &keliling)
	default:
		fmt.Println("jenis tidak valid")
	}

	fmt.Println("Luas : ", luas)
	fmt.Println("Keliling : ", keliling)
}
