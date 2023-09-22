package mypackage

import (
	"fmt"
)

func LearnVariable() {
	// * variabel menggunakan tipe data
	var name string = "danar"
	fmt.Println(name)

	// * variabel tanpa tipe data
	age := 20
	fmt.Println(age)

	// * multi variables
	var first, second, third string = "first", "second", "third"
	fourth, fifth, sixth := "fourth", "fifth", "sixth"
	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)

	// * variabel underscore
	_ = "underscore" // konsepnyta seperti kerangjang sampah

	// * variabel pointer dengan keyword new
	address := new(string)
	address = &name
	fmt.Println(address)
	fmt.Println(&address)
	fmt.Println(*address)

	// * nilain nil adalah kosong. berbeda dengan zero value

	// * konstanta
	const (
		pi             = 3.14
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)
	fmt.Println(isToday)

	// * penkondisian
	var point = 8
	if percent := point / 2; percent == 10 {
		fmt.Println("perfect")
	} else if point < 5 {
		fmt.Println("awesome")
	}

	// * perulangan
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	// atau
	var buah = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range buah { // i wajib digunakan
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	// atau
	for _, fruit := range buah {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// * array
	var fruits = [...]string{"apple", "grape", "banana"}
	fmt.Println(fruits)

	// * slice versi reference dari array
	var fruitsSlice = []string{"apple", "grape", "banana"}
	var newFruitsSlice = fruitsSlice[0:2]
	newFruitsSlice[0] = "papaya"
	fmt.Println(fruitsSlice) // [papaya grape banana]

	// * map
	var chicken map[string]int
	chicken = map[string]int{} // inisialisasi map agar tidak nil
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])
}
