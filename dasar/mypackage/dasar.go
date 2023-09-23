package mypackage

import (
	"fmt"
	"strings"
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

	// * kombinasi slice dan map
	var chickenCombine = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "female"},
	}
	for _, chicken := range chickenCombine {
		fmt.Print(chicken["gender"], " ", chicken["name"], "|")
	}
	fmt.Print("\n")

	// * fungsi (closure)
	adding := func(a, b int) (int, float32) {
		return a + b, float32(a) / float32(b)
	}
	fmt.Println(adding(1, 2))

	// * closure return fungsi
	many, max := findMax([]int{1, 2, 3, 4, 5, 1, 2, 10, 10}, 5)
	fmt.Println(many, max())

	// * closure sebagai parameter
	var data = []string{"wick", "jason", "ethan", "danar", "falal", "john"}

	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data asli \t\t:", data)                     // data asli : [wick jason ethan]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)  // filter ada huruf "o" : [jason]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5) // filter jumlah huruf "5" : [jason ethan]

	// * pointer
	var numberA int = 4
	var numberB *int = &numberA                     // nilai pointer numberA atau address
	numberA = 5                                     // ubah nilai numberA
	fmt.Println("numberA (value)   :", numberA)     // 4
	fmt.Println("numberA (address) :", &numberA)    // 0xc0000140a8
	fmt.Println("numberB (value)   :", *numberB)    // 4
	fmt.Println("numberB (address) :", numberB)     // 0xc0000140a8
	fmt.Println("numberB (true address)", &numberB) // 0xc00000e028

	// * struct
	p := person{name: "danar", age: 20}
	fmt.Println(p.getName()) // danar

}

// * struct
type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return p.name
}

// * closure return fungsi
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

// * closure sebagai parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
