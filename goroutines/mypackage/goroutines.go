package mypackage

import (
	"fmt"
	"runtime"
	"time"
)

// * goroutines
func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func LearnGoroutines() {
	// Membuat dua goroutine
	runtime.GOMAXPROCS(2)
	go bar()
	go foo()

	// Menunggu agar kedua goroutine selesai
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Selesai")
}
