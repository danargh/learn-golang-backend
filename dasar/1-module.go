package main

import (
	"fmt"

	say_hello "github.com/danargh/my-golang-module-learning"
)

func main() {
	var hello string = say_hello.SayHello("danar")
	fmt.Println(hello)
}
