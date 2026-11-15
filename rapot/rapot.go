package main

import "fmt"

func NH(x int) string {

	if x > 80 {
		return "A"
	} else if x > 70 && x <= 80 {
		return "AB"
	} else if x > 65 && x <= 70 {
		return "B"
	} else if x > 60 && x <= 65 {
		return "BC"
	} else if x > 50 && x <= 60 {
		return "C"
	} else if x > 40 && x <= 50 {
		return "D"
	} else {
		return "E"
	}
}

func rata(total, n int) float64 {
	return float64(total) / float64(n)
}

func main() {

	var n int
	var b int
	var name string

	c := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&name, &b)
		fmt.Println(name, b, NH(b))
		c += b

	}
	fmt.Printf("Rata-rata: %.2f\n", rata(c, n))

}
