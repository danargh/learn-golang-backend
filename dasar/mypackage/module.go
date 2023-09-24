// mypackage/module.go
package mypackage

import (
	f "fmt"

	say_hello "github.com/danargh/my-golang-module-learning"
)

func ConsumeModule() {
	var hello string = say_hello.SayHello("danar")
	f.Println(hello)

	// * exported module
	// fungsi, struct, propety yang memiliki huruf awal kapital merupakan exported / public

}
