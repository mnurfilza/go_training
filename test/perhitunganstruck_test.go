package test

import (
	"Training_go/model"
	"fmt"
	"testing"
)

var dataBalok = []struct {
	inputBalok  model.Balok
	Kelilingnya int
	Volumenya   int
	Luasnya     int
}{
	{inputBalok: model.Balok{Panjang: 2, Lebar: 3, Tinggi: 4}, Kelilingnya: 36, Volumenya: 24, Luasnya: 52},
	{inputBalok: model.Balok{Panjang: 2, Lebar: 2, Tinggi: 2}, Kelilingnya: 24, Volumenya: 8, Luasnya: 24},
	{inputBalok: model.Balok{Panjang: 5, Lebar: 4, Tinggi: 3}, Kelilingnya: 48, Volumenya: 60, Luasnya: 94},
}

var dataSegitiga = []struct {
	inputSegitiga model.Segitiga
	Kelilingnya   int
	Tingginya     float64
	Luasnya       float64
}{
	{inputSegitiga: model.Segitiga{Alas: 12, Sisi: 10}, Kelilingnya: 32, Tingginya: 8, Luasnya: 48},
	{inputSegitiga: model.Segitiga{Alas: 6, Sisi: 5}, Kelilingnya: 16, Tingginya: 4, Luasnya: 12},
	{inputSegitiga: model.Segitiga{Alas: 60, Sisi: 50}, Kelilingnya: 160, Tingginya: 40, Luasnya: 1200},
}

func TestPerhitunganStruct(t *testing.T) {
	t.Run("Tes untuk string", func(t *testing.T) {
		for _, item := range dataBalok {
			dapatKelilingnya := item.inputBalok.Keliling()
			if dapatKelilingnya != item.Kelilingnya {
				t.Fatalf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Kelilingnya= %v, mau Kelilingnya= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatKelilingnya, item.Kelilingnya)
			}
			fmt.Printf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Kelilingnya= %v, mau Kelilingnya= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatKelilingnya, item.Kelilingnya)
		}
	})

	t.Run("Tes untuk string", func(t *testing.T) {
		for _, item := range dataBalok {
			dapatVolumenya := item.inputBalok.Volume()
			if dapatVolumenya != item.Volumenya {
				t.Fatalf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Kelilingnya= %v, mau Volume= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatVolumenya, item.Volumenya)
			}
			fmt.Printf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Kelilingnya= %v, mau Volume= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatVolumenya, item.Volumenya)
		}
	})
	t.Run("Tes untuk string", func(t *testing.T) {
		for _, item := range dataBalok {
			dapatLuasnya := item.inputBalok.Luas()
			if dapatLuasnya != item.Luasnya {
				t.Fatalf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Luasnya= %v, mau Luasnya= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatLuasnya, item.Luasnya)
			}
			fmt.Printf("Panjang= %v, Lebar= %v, Tinggi = %v, maka Luasnya= %v, mau Luasnya= %v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, dapatLuasnya, item.Luasnya)
		}
	})
}

func TestPerhitunganSegitiga(t *testing.T) {
	t.Run("Tes Keliling Segitiga ", func(t *testing.T) {
		for _, item := range dataSegitiga {
			dapatKelilingnya := item.inputSegitiga.KelilingSegitiga()
			if dapatKelilingnya != item.Kelilingnya {
				t.Fatalf("Alasnya : %v, Sisinya : %v , Tingginya : %v, Kelilingnya : %v , Maunya Keliling : %v\n", item.inputSegitiga.Alas, item.Tingginya, item.inputSegitiga.Sisi, item.Kelilingnya, dapatKelilingnya)
			}
			fmt.Printf("Alasnya : %v, Sisinya : %v , Tingginya : %v, Kelilingnya : %v , Maunya Keliling : %v\n", item.inputSegitiga.Alas, item.Tingginya, item.inputSegitiga.Sisi, item.Kelilingnya, dapatKelilingnya)
		}
	})
	t.Run("Tes Tinggi Segitiga ", func(t *testing.T) {
		for _, item := range dataSegitiga {
			dapatTingginya := item.inputSegitiga.TinggiSegitiga()
			if dapatTingginya != item.Tingginya {
				t.Fatalf(" Tingginya : %v,  Maunya Tinggi : %v\n", item.Tingginya, dapatTingginya)
			}
			fmt.Printf(" Tingginya : %v,  Maunya Tinggi : %v\n", item.Tingginya, dapatTingginya)
		}
	})
	t.Run("Tes Luas Segitiga ", func(t *testing.T) {
		for _, item := range dataSegitiga {
			dapatLuasnya := item.inputSegitiga.LuasSegitiga()
			if dapatLuasnya != item.Luasnya {
				t.Fatalf("Alasnya : %v, Sisinya : %v , Tingginya : %v, Kelilingnya : %v , Maunya Keliling : %v\n", item.inputSegitiga.Alas, item.Tingginya, item.inputSegitiga.Sisi, item.Luasnya, dapatLuasnya)
			}
			fmt.Printf("Alasnya : %v, Sisinya : %v , Tingginya : %v, Kelilingnya : %v , Maunya Keliling : %v\n", item.inputSegitiga.Alas, item.Tingginya, item.inputSegitiga.Sisi, item.Luasnya, dapatLuasnya)
		}
	})
}
