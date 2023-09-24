package main

import (
	"fmt"
	"learn-golang/goroutines/mypackage"
)

func main() {
	// learn goroutines
	fmt.Println("learn goroutines")
	mypackage.LearnGoroutines()

	// learn channel
	fmt.Println("\nlearn channel")
	mypackage.LearnChannel()
}
