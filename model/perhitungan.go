package model

import (
	"fmt"
	"time"
)

func Penjumlahan(a, b int) int {
	return a + b

}

func Perkalian(a, b int) int {
	return a * b
}

func Pembagian(a, b int) int {
	return a / b
}

func Pengurangan(a, b int) int {
	return a - b
}

func Pemangkatan(a, b int) int {
	hasil := 1
	for i := 0; i < b; i++ {
		hasil = hasil * a
	}
	return hasil
}

func Max(m []int) int {
	if len(m) == 0 {
		return 0
	}
	maksimum := m[0]
	for _, nilai := range m {
		if nilai > maksimum {
			maksimum = nilai
		}
	}
	return maksimum
}

func Min(m []int) int {
	if len(m) == 0 {
		return 0
	}
	minimum := m[0]
	for _, nilai := range m {
		if nilai < minimum {
			minimum = nilai
		}
	}
	return minimum
}

func Fac(a int) int {

	hasil := 1
	for i := 1; i <= a; i++ {
		hasil = hasil * i
	}
	return hasil

}

func Ganjil(n int) (bilanganGanjil, bilanganGenap []int) {

	for i := 0; i < n; i++ {
		hasil := (i * 2)
		bilanganGanjil = append(bilanganGanjil, hasil+1)
		bilanganGenap = append(bilanganGenap, hasil)
	}
	return
}
func TanggalLahir(tanggalLahir time.Time) string {
	tahun, month, day := tanggalLahir.Date()
	bulan := int(month)
	dateBorn := fmt.Sprintf("%v-%v-%v", tahun, bulan, day)
	return dateBorn

}
func TambahString(a int, s string) string {
	var hasil string

	for i := 0; i < a; i++ {
		hasil += s
	}
	return hasil
}
