package main

import (
	f "fmt"
	"learn-golang/dasar/mypackage"
)

func main() {
	// learn golang module :
	f.Println("\n>>> Learn Golang Module : \n")
	mypackage.ConsumeModule()

	// learn variabel :
	f.Println("\n>>> Learn Variabel : \n")
	mypackage.LearnVariable()

	// learn interface
	f.Println("\n>>> Learn Interface : \n")
	mypackage.LearnInterface()
}
