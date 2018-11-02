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
