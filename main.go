package main

import (
	"fmt"
)

// 1. menghitung luas permukaan tabung
func main() {
	var phi float32
	phi = 3.14
	var jari float32
	jari = 3
	var tinggi float32
	tinggi = 9

	var rumus float32
	rumus = 2*phi*jari*jari + tinggi
	fmt.Println(rumus)

	// 2. grade nilai mahasiswa
	nilai := 54
	var grade string

	if nilai < 55 {
		grade = "TIDAK LULUS"
	} else if nilai >= 55 {
		grade = "D"
	} else if nilai >= 65 {
		grade = "C"
	} else if nilai >= 75 {
		grade = "B"
	} else if nilai >= 85 {
		grade = "A"
	} else {
		grade = "null"
	}

	fmt.Println(grade)

	// 3. urutan penjumlahan
	var urutanNilai [6]int
	urutanNilai[0] = 1
	urutanNilai[1] = 2
	urutanNilai[2] = 3
	urutanNilai[3] = 4
	urutanNilai[4] = 5
	urutanNilai[5] = 6
	var sum int
	sum = urutanNilai[0] + urutanNilai[1] + urutanNilai[2] + urutanNilai[3] + urutanNilai[4] + urutanNilai[5]

	fmt.Println(sum)

	// 4. bintang
	for b := 1; b <= 5; b++ {
		for j := 1; j <= b; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
	return
}
