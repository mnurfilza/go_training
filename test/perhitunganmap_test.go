package test

import (
	"Training_go/model"
	"fmt"
	"testing"
)

func TestPerhitunganMap(t *testing.T) {
	t.Run("test untuk fungsi penjumlahan ", func(t *testing.T) {
		var testPenjumlahan = []struct {
			s           string
			P           int
			L           int
			T           int
			hasilMaunya interface{}
		}{
			{s: "Keliling", P: 2, L: 2, T: 2, hasilMaunya: 24},
			{s: "Luas", P: 2, L: 3, T: 4, hasilMaunya: 52},
			{s: "Volume", P: 5, L: 4, T: 3, hasilMaunya: 60},
		}

		for _, input := range testPenjumlahan {
			dapatHasil, errorDapatnya := model.PerhitunganMapBalok(input.s, input.P, input.L, input.T)
			if dapatHasil != input.hasilMaunya || errorDapatnya != nil {
				t.Fatalf("Hasil Maunya %v, Hasil Dapatnya %v, errornya %v\n", input.hasilMaunya, dapatHasil, errorDapatnya)
			}
			fmt.Printf("Hasil Maunya %v, Hasil Dapatnya %v, errornya %v\n", input.hasilMaunya, dapatHasil, errorDapatnya)
		}

	})
	t.Run("test untuk fungsi penjumlahan ", func(t *testing.T) {
		var testSegitiga = []struct {
			s           string
			Alasnya     int
			Sisinya     int
			hasilMaunya interface{}
		}{
			{s: "Keliling", Alasnya: 12, Sisinya: 10, hasilMaunya: 32},
			{s: "Luas", Alasnya: 12, Sisinya: 10, hasilMaunya: 48.0},
			{s: "Tinggi", Alasnya: 12, Sisinya: 10, hasilMaunya: 8.0},
		}

		for _, input := range testSegitiga {
			dapatHasil, errorDapatnya := model.PerhitunganMapSegitiga(input.s, input.Alasnya, input.Sisinya)
			if dapatHasil != input.hasilMaunya || errorDapatnya != nil {
				t.Fatalf("Hasil Maunya %v, Hasil Dapatnya %v, errornya %v\n", input.hasilMaunya, dapatHasil, errorDapatnya)
			}
			fmt.Printf("Hasil Maunya %v, Hasil Dapatnya %v, errornya %v\n", input.hasilMaunya, dapatHasil, errorDapatnya)
		}

	})

}
