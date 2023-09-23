// mypackage/module.go
package mypackage

import (
	"fmt"

	say_hello "github.com/danargh/my-golang-module-learning"
)

func ConsumeModule() {
	var hello string = say_hello.SayHello("danar")
	fmt.Println(hello)
}
