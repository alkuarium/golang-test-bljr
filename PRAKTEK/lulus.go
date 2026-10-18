package main

import "fmt"

func main() {

	var n int

	for {
		fmt.Scan(&n)

		if n >= 60 {
			fmt.Println("Anda lulus ujian")
		} else {
			fmt.Println("Anda tidak lulus ujian")
		}
		if n == -1 {
			fmt.Println("program selesai")
			break
		}
	}
}
