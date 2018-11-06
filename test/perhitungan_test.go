package test

import (
	"Training_go/model"
	"fmt"
	"testing"
	"time"
)

func TestPerhitungan(t *testing.T) {
	t.Run("test untuk fungsi penjumlahan ", func(t *testing.T) {
		var testPenjumlahan = []struct {
			a           int
			b           int
			hasilMaunya int
		}{
			{a: 1, b: 1, hasilMaunya: 2},
			{a: 1, b: 3, hasilMaunya: 4},
			{a: 2, b: 3, hasilMaunya: 5},
			{a: 7, b: 3, hasilMaunya: 10},
			{a: -1, b: -4, hasilMaunya: -5},
		}

		for _, input := range testPenjumlahan {
			dapatHasil := model.Penjumlahan(input.a, input.b)
			if dapatHasil != input.hasilMaunya {
				t.Fatalf("%v + %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.hasilMaunya)
			}
			fmt.Printf("%v + %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.hasilMaunya)
		}

	})
	t.Run("test untuk fungsi perkalian ", func(t *testing.T) {
		var testPerkalian = []struct {
			a   int
			b   int
			got int
		}{
			{a: 1, b: 1, got: 1},
			{a: 1, b: 3, got: 3},
			{a: 2, b: 3, got: 6},
			{a: 7, b: 3, got: 21},
		}

		for _, input := range testPerkalian {
			dapatHasil := model.Perkalian(input.a, input.b)
			if dapatHasil != input.got {
				t.Fatalf("%v * %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.got)
			}
			fmt.Printf("%v * %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.got)
		}
	})
	t.Run("test untuk fungsi pemangkatan ", func(t *testing.T) {
		var testPemangkatan = []struct {
			a   int
			b   int
			got int
		}{
			{a: 2, b: 2, got: 4},
			{a: 3, b: 3, got: 27},
			{a: 5, b: 2, got: 25},
			{a: 11, b: 2, got: 121},
		}
		for _, input := range testPemangkatan {
			dapatHasil := model.Pemangkatan(input.a, input.b)
			if dapatHasil != input.got {
				t.Fatalf("%v * %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.got)
			}
			fmt.Printf("%v * %v = Got : %v , Want = %v\n", input.a, input.b, dapatHasil, input.got)
		}
	})

	t.Run("Test untuk Fungsi max ", func(t *testing.T) {
		var testMax = []struct {
			a           []int
			hasilMaunya int
		}{
			{a: []int{1, 4, 6, 8, 7}, hasilMaunya: 8},
			{a: []int{2, 4, 5, 12, 7}, hasilMaunya: 12},
		}

		for _, input := range testMax {
			dapatHasil := model.Max(input.a)
			if dapatHasil != input.hasilMaunya {
				t.Fatalf("%v Got : %v , want : %v\n ", input.a, dapatHasil, input.hasilMaunya)
			}
			fmt.Printf("%v Got : %v , want : %v\n ", input.a, dapatHasil, input.hasilMaunya)
		}
	})
	t.Run("Test untuk Fungsi max ", func(t *testing.T) {
		var testMin = []struct {
			a           []int
			hasilMaunya int
		}{
			{a: []int{1, 4, 6, 8, 7}, hasilMaunya: 1},
			{a: []int{2, 4, 5, 12, 7}, hasilMaunya: 2},
		}

		for _, input := range testMin {
			dapatHasil := model.Min(input.a)
			if dapatHasil != input.hasilMaunya {
				t.Fatalf("%v Got : %v , want : %v\n ", input.a, dapatHasil, input.hasilMaunya)
			}
			fmt.Printf("%v Got : %v , want : %v\n ", input.a, dapatHasil, input.hasilMaunya)
		}
	})
	t.Run("Test Untuk Mengetahui Fungsi Faktorial : ", func(t *testing.T) {
		var testFac = []struct {
			a   int
			got int
		}{
			{a: 3, got: 6},
			{a: 4, got: 24},
		}
		for _, input := range testFac {
			dapatHasil := model.Fac(input.a)
			if dapatHasil != input.got {
				t.Fatalf("input %v got : %v want : %v\n", input.a, dapatHasil, input.got)
			}
			fmt.Printf("input %v got : %v want : %v\n", input.a, dapatHasil, input.got)
		}

	})
	t.Run("Test Untuk Mengetahui Fungsi Tanggal : ", func(t *testing.T) {
		var testDate = []struct {
			date time.Time
			got  string
		}{
			{date: time.Date(1995, 4, 5, 0, 0, 0, 0, time.UTC), got: "1995-4-5"},
			{date: time.Date(1995, 5, 3, 0, 0, 0, 0, time.UTC), got: "1995-5-3"},
			{date: time.Date(1995, 6, 5, 0, 0, 0, 0, time.UTC), got: "1995-6-5"},
			{date: time.Date(1995, 7, 5, 0, 0, 0, 0, time.UTC), got: "1995-7-5"},
			{date: time.Date(1995, 9, 5, 0, 0, 0, 0, time.UTC), got: "1995-9-5"},
			{date: time.Date(1995, 11, 5, 0, 0, 0, 0, time.UTC), got: "1995-11-5"},
		}
		for _, input := range testDate {
			dapatHasil := model.TanggalLahir(input.date)
			if dapatHasil != input.got {
				t.Fatalf("time.Time nya  %v got : %v want : %v\n", input.date, dapatHasil, input.got)
			}
			fmt.Printf("time.Tme nya %v got : %v want : %v\n", input.date, dapatHasil, input.got)
		}

	})

	t.Run("Test Untuk Mengetahui Bilangan Ganjil dan Genap ", func(t *testing.T) {
		var tesBil = []struct {
			n              int
			bilanganGanjil []int
			bilanganGenap  []int
		}{
			{n: 3, bilanganGanjil: []int{1, 3, 5}, bilanganGenap: []int{0, 2, 4}},
			{n: 4, bilanganGanjil: []int{1, 3, 5, 7}, bilanganGenap: []int{0, 2, 4, 6}},
		}
		for _, input := range tesBil {
			hasilDapatnyaGanjil, hasilDapatnyaGenap := model.Ganjil(input.n)
			if len(hasilDapatnyaGanjil) != len(input.bilanganGanjil) && len(hasilDapatnyaGenap) != len(input.bilanganGenap) {
				t.Fatalf("Length Maunya ganjil %v, Length Dapatnya ganjil %v, length maunya genap %v, length dapatnya genap %v\n", len(hasilDapatnyaGanjil), len(input.bilanganGanjil), len(hasilDapatnyaGenap), len(input.bilanganGenap))
			}
			for index, value := range hasilDapatnyaGanjil {
				if value != input.bilanganGanjil[index] {
					t.Fatalf("Nilai Ganjilnya : %v, dapat Ganjilnya : %v, Mau Ganjilnya %v\n", input.n, value, input.bilanganGanjil[index])
				}
				fmt.Printf("Nilai Ganjilnya : %v, dapat Ganjilnya : %v, Mau Ganjilnya %v\n", input.n, value, input.bilanganGanjil[index])

			}
			for index, value := range hasilDapatnyaGenap {
				if value != input.bilanganGenap[index] {
					t.Fatalf("Nilai Genapnya : %v, dapat Genapnya : %v, Mau Ganjilnya %v\n", input.n, value, input.bilanganGenap)

				}
				fmt.Printf("Nilai Genapnya : %v, dapat Genapnya : %v, Mau Ganjilnya %v\n", input.n, value, input.bilanganGenap)
			}
		}

	})
	t.Run("test String", func(t *testing.T) {
		var tesString = []struct {
			a            int
			s            string
			maunyaString string
		}{
			{a: 1, s: "a", maunyaString: "a"},
			{a: 2, s: "ab", maunyaString: "abab"},
			{a: 3, s: "a", maunyaString: "aaa"},
			{a: 4, s: "abc", maunyaString: "abcabcabcabc"},
		}
		for _, input := range tesString {
			dapatString := model.TambahString(input.a, input.s)
			if dapatString != input.maunyaString {
				t.Fatalf("angkanya %v, textnya %v, dapat Stringnya %v , Maunya %v", input.a, input.s, dapatString, input.maunyaString)
			}
			fmt.Printf("angkanya %v, textnya %v, dapat Stringnya %v , Maunya %v\n", input.a, input.s, dapatString, input.maunyaString)
		}
	})

}
