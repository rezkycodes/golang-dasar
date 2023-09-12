package main

import "fmt"

func main() {
	sayHelloTo("Rezky", "Budihartono")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
