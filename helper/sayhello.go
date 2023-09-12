package helper

import "fmt"

var version = 1
var Application = "belajar golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string)  {
	fmt.Println("Goodbye",name)
}