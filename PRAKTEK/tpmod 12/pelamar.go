package main

import "fmt"

const NMAX = 100

type plmr struct {
	id, name  string
	testGrade float64
	durasites int
}
type plmr2 [NMAX]plmr

func lolos(a plmr2, n int, batas float64, b *plmr2, m *int) {
	*m = 0
	for i := 0; i < n; i++ {
		if a[i].testGrade >= batas {
			b[*m] = a[i]
			*m++
		}
	}
}

func durasi(b plmr2, m int) float64 {
	if m == 0 {
		return 0.0
	}

	totalDurasi := 0
	for i := 0; i < m; i++ {
		totalDurasi += b[i].durasites
	}
	return float64(totalDurasi) / float64(m)
}

func main() {
	var a, b plmr2
	var n, m int
	var batas float64
	var id string

	n = 0
	fmt.Scan(&id)
	for id != "END" && n < NMAX {
		a[n].id = id
		fmt.Scan(&a[n].name, &a[n].testGrade, &a[n].durasites)
		n++
		fmt.Scan(&id)
	}

	fmt.Scan(&batas)

	lolos(a, n, batas, &b, &m)

	for i := 0; i < n; i++ {
		fmt.Println(a[i].id, a[i].name, a[i].testGrade, a[i].durasites)
	}
	fmt.Printf("Pelamar lolos (nilai >= %02.f)\n", batas)
	if m == 0 {
		fmt.Println("Tidak ada yang lolos")
	}
	for i := 0; i < m; i++ {
		fmt.Println(b[i].id, b[i].name, b[i].testGrade, b[i].durasites)
	}
	fmt.Println("jumlah lolos = ", m)
	fmt.Println("rata-rata durasi : ", durasi(b, m), "menit")
}
