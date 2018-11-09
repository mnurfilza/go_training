package main

import (
	"fmt"
	"go_training/model"
	"time"
)

var c, d = 20, 10
var e, f = 11, 2
var g = []int{4, 5, 6, 17, 10}
var w = 4
var tanggalLahir = time.Date(1955, 5, 7, 0, 0, 0, 0, time.UTC)
var s = "abc"
var a = 2

var S = model.Segitiga{Alas: 6, Sisi: 5}
var x = map[string]model.Mahasiswa{
	"Nama": {Universitas: "UI", SMA: "Sman45", SMP: "smpn1", SD: "sd2"},
	"aaan": {Universitas: "UA", SMA: "Sman46", SMP: "smpn2", SD: "sd3"},
}

func main() {
	fmt.Print(model.PerhitunganMapMahasiswa(x))
}

func loop() {
	for x := 1; x <= 5; x++ {
		for i := 1; i <= 5; i++ {
			if i == 1 || i == 5 {
				fmt.Print("0")
			} else {
				if x == 1 || x == 5 {
					fmt.Print("0")
				} else {
					fmt.Print(" ")
				}

			}
		}
		fmt.Println()
	}
}

func switching() {
	angka := 0
	fmt.Print("Masukan Angka 1-5 : ")
	fmt.Scanf("%d", &angka)
	switch angka {
	case 1:
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	case 2:
		fmt.Print("Menampilkan Angkka 2")
	case 3:
		fmt.Print("Menampilkan Angkka 3")
	case 4:
		loop()
	case 5:
		jumlah := model.Penjumlahan(c, d)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", jumlah)
	case 6:
		jumlah := model.Perkalian(c, d)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", jumlah)
	case 7:
		jumlah := model.Pengurangan(c, d)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", jumlah)
	case 8:
		jumlah := model.Pembagian(c, d)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", jumlah)
	case 9:
		pangkat := model.Pemangkatan(e, f)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", pangkat)
	case 10:
		r := 5

		fac := model.Fac(r)
		fmt.Printf("Hasil Dari Fungsi Penjumlahan = %d", fac)
	default:
		fmt.Print("Angka tidak ada")
	}

}
