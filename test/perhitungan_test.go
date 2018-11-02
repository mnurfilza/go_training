package test

import (
	"Training_go/model"
	"fmt"
	"testing"
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
}
