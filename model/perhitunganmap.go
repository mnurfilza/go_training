package model

import (
	"errors"
	"fmt"
	"strings"
)

var buatMap = make(map[string]interface{})

func PerhitunganMapBalok(s string, panjang int, lebar int, tinggi int) (interface{}, error) {
	balok := Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}
	buatMap[s] = balok

	switch s {
	case "Keliling":
		buatMap["Keliling"] = balok.Keliling()
		return buatMap["Keliling"], nil
	case "Luas":
		buatMap["Luas"] = balok.Luas()
		return buatMap["Luas"], nil
	case "Volume":
		buatMap["Volume"] = balok.Volume()
		return buatMap["Volume"], nil
	default:
		return nil, errors.New("Key tidak ada!")
	}
}

func PerhitunganMapSegitiga(s string, alas int, sisi int) (interface{}, error) {
	segitiga := Segitiga{Alas: alas, Sisi: sisi}
	buatMap[s] = segitiga

	switch s {
	case "Keliling":
		buatMap["Keliling"] = segitiga.KelilingSegitiga()
		return buatMap["Keliling"], nil
	case "Tinggi":
		buatMap["Tinggi"] = segitiga.TinggiSegitiga()
		return buatMap["Tinggi"], nil
	case "Luas":
		buatMap["Luas"] = segitiga.LuasSegitiga()
		return buatMap["Luas"], nil
	default:
		return nil, errors.New("Key tidak ada!")
	}
}

func PerhitunganMapMahasiswa(bio map[string]Mahasiswa) string {
	var hasil string
	for index, value := range bio {
		hasil += fmt.Sprintf("%v %v", index, value)
	}
	return hasil
}

func WordCount(s string) map[string]int {
	Map := make(map[string]int)
	theSplit := strings.Fields(s)
	for _, nilai := range theSplit {
		Map[nilai] = Map[nilai] + 1
	}
	return Map
}
