package model

import (
	"fmt"
	"math"
)

type Balok struct {
	Panjang int
	Lebar   int
	Tinggi  int
}

type Segitiga struct {
	Alas int
	Sisi int
}

type Mahasiswa struct {
	Universitas string
	SMA         string
	SD          string
	SMP         string
}

func (b *Balok) Keliling() int {
	return 4 * (b.Panjang + b.Lebar + b.Tinggi)
}
func (b *Balok) Volume() int {
	return b.Panjang * b.Lebar * b.Tinggi
}
func (b *Balok) Luas() int {
	return 2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi))
}

func (s *Segitiga) KelilingSegitiga() int {
	return s.Sisi + s.Sisi + s.Alas
}

func (s *Segitiga) TinggiSegitiga() float64 {
	tinggi := (s.Sisi * s.Sisi) - (s.Alas/2)*(s.Alas/2)
	t := float64(tinggi)
	hasil := math.Sqrt(t)
	return hasil
}

func (s *Segitiga) LuasSegitiga() float64 {
	tinggi := Segitiga{Alas: s.Alas, Sisi: s.Sisi}
	return 0.5 * float64(s.Alas) * tinggi.TinggiSegitiga()
}

func (bio Mahasiswa) String() string {
	return fmt.Sprintf("Saya kuliah di %v, SMA saya di %v, SMP saya di %v, SD saya di %v", bio.Universitas, bio.SMA, bio.SMP, bio.SD)
}
