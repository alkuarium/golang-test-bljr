package main 

import "fmt"

func Array() {

	// SLICE
	fmt.Println("---SLICE---")

	buah := []string{"Mangga", "Apel", "Jeruk", "Pisang", "Semangka"}

	a := buah[:1]
	b := buah[2:5]
	buah[0] = "Durian"
	addbuah := append(buah, "Nanas", "Melon")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(addbuah)


	//ARRAY BIASA
	data := [5]int{1, 2, 3, 4, 5}
	fmt.Println(data)

	dataBuah := [5]string{"Mangga", "Apel", "Jeruk", "Pisang", "Semangka"}
	fmt.Println(dataBuah)
	aku := len(dataBuah)
	fmt.Println(aku)



	//MAP
	fmt.Println("---MAP---")

	car := map[string]string{
		"Car1":  "Toyota",
		"Car2":  "Honda",
		"Car3":  "Suzuki",
		"Car4":  "Mitsubishi",
		"Car5":  "Daihatsu",
	}


	car["Car3"] = "yamaha"
	delete(car, "Car5")
	hitngCAR := len(car)

	fmt.Println(car)
	fmt.Println(hitngCAR)
	fmt.Println(car["Car4"])
}
