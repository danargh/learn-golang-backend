// testing

package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rudi")
	if result != "Hello Rudi" {
		// error
		t.Fatal("Error")
	}
	fmt.Println("Success")
}
