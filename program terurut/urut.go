package main

import "fmt"

func tukar(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func urut(a, b, c *int) {
	if *a > *b {
		tukar(a, b)
	}
	if *b > *c {
		tukar(b, c)
	}
	if *a > *b {
		tukar(a, b)
	}
}
func main() {
	var d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)
	fmt.Println("sebelum :", d1, d2, d3)
	urut(&d1, &d2, &d3)
	fmt.Println("sesudah :", d1, d2, d3)

}
