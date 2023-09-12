package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {

	nameFiltered := filter(name)

	fmt.Println("Hello", nameFiltered)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func upperFilter(name string) string {
	return strings.ToUpper(name)
}

func main() {
	sayHelloWithFilter("Rezky", upperFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
