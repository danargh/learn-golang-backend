package main

import "fmt"

func main() {
	fmt.Println("Variabel")
	// deklarasi variabel dengan banyak cara
	nama := "danar"
	lastName := new(string)
	*lastName = "ghulamsyah"
	fmt.Printf("halo %s %s \n", nama, *lastName)

	// deklarasi multi variabel
	var name, age = "danar", 23
	fmt.Printf("halo %s umur %d \n", name, age)

	// tipe data
	fmt.Println("Tipe Data", "dasar")
	var x int = 10
	fmt.Printf("tipe data x adalah %T \n", x)
	fmt.Printf("nilai variabel x adalah %d \n", x)

	// konstanta
	const (
		phi = 3.14
	)
	fmt.Print("nilai phi : ", phi, "\n")

	// array
	var arr = [...]int{1, 2, 3}
	fmt.Println("ini array : ", arr)
	fmt.Println("arr length : ", len(arr))

	// slice
	iniSlice := []int{1, 2, 3}
	fmt.Println("ini slice : ", iniSlice)

	// map
	fmt.Println("Map")
	var chicken map[int]string
	chicken = map[int]string{
		1: "ayam jago",
	}
	chicken[1] = "ayam kate"
	fmt.Println(chicken[1])

	// pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	// struct
	type student struct {
		name  string
		grade int
	}

	// struct implementation
	var s1 student
	s1.name = "danar"
	s1.grade = 2

	// inisialisasi object struct
	var s2 = student{}
	s2.name = "danar"
	s2.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s2.grade)
}
