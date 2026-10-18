package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	Bawal := 10000000.0

	for i := 2; i <=n; i++{
		Bawal = Bawal * 1.06
	}
		fmt.Printf("%.0f\n", Bawal)
}
