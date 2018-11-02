package model

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
