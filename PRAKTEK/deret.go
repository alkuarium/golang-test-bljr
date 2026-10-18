package main

import "fmt"

func main() {
	var n int
	tambah := 0 //bertugas sebagai penampung hasil penjumlahan
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		tambah += i //menampung hasil lalu menambahkannya 
		fmt.Print(i, " + ")
	}
	fmt.Println( tambah)
}
