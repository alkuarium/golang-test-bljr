package main

import "fmt"

func main() {
	var n, nilai int
	var nama string
	fmt.Scan(&n)

	var namaTerendah1, namaTerendah2 string
	nilaiTerendah1 := 999
	nilaiTerendah2 := 999

	i := 1
	for i <= n {
		fmt.Scan(&nama)
		fmt.Scan(&nilai)

		if nilai < nilaiTerendah1 {
			namaTerendah1 = nama
			nilaiTerendah1 = nilai
		} else if nilai < nilaiTerendah2 {
			nilaiTerendah2 = nilai
			namaTerendah2 = nama
		}


		i++
	}
	fmt.Println(namaTerendah1,nilaiTerendah1, "&", namaTerendah2,nilaiTerendah2)

}
