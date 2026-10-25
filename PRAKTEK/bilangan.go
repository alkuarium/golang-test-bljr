package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n == 0 {
		fmt.Print("nol")
		return
	}
	if n%2 == 0 {
		fmt.Print("genap, ")
		if n > 0 {
			fmt.Print("positif")
		} else if n < 0 {
			fmt.Print("negatif")
		}
	} else {
		fmt.Print("ganjil, ")
		if n > 0 {
			fmt.Print("positif")
		} else if n < 0 {
			fmt.Print("negatif")
		}
	}
}
