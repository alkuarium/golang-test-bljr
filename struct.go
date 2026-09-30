package main

import "fmt"

//contoh struct
type orang struct {
	nama    string
	umur    int
	id      int
	bio     string
	menikah bool
}

//contoh interface
type bentuk interface {
	agk() uint
}

type kotak struct {
	width uint
}

type segitiga struct {
	a uint
	b uint
	c uint
}

func (se segitiga) agk() uint {
	return se.a * se.b * se.c
}

func (ko kotak) agk() uint {
	return 4 * ko.width
}


func main() {
	//contoh kalau kosong
	// orang1 := orang{}
	// fmt.Println(orang1)

	// //contoh ada isinya
	// orang2 := orang{
	// 	nama: "AL",
	// 	umur:    14,
	// 	id:      65,
	// 	bio:     "anjay",
	// 	menikah: true}
	// 	orang2.bio = "walla"
	// fmt.Println(orang2)

	//contoh jika ingin memasukan nilai lewat terminal
	// var p orang
	// fmt.Scan(&p.nama, &p.umur, &p.id, &p.bio, &p.menikah)

	// fmt.Println(p)


	//contoh untuk pemanggilan interface
	var bentuk1 bentuk =segitiga{4,5,6}
	fmt.Println(bentuk1)
	fmt.Println(bentuk1.agk())

	var bentuk2 bentuk =kotak{4}
	fmt.Println(bentuk2.agk())
}
