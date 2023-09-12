package main

import "fmt"

func main() {
	sayHello("Rezky", "Budihartono")
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
