package mypackage

import (
	f "fmt"
)

// * interface
type hitung interface {
	luas() float32
	keliling() float32
}
type persegi struct {
	sisi float32
}

func (p persegi) luas() float32 {
	return p.sisi * p.sisi
}
func (p persegi) keliling() float32 {
	return 4 * p.sisi
}

func LearnInterface() {

	// * interface
	// apabila hitung diberi struct selain persegi maka interface akan tau
	var bangunDatar hitung = persegi{10.0}
	f.Println("Luas", bangunDatar.luas())
	f.Println("Keliling", bangunDatar.keliling())

}
