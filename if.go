package main

import "fmt"

func main() {
	name := "rezky"

	if name == "rezkys" {
		fmt.Println(name)
	} else if name == "baron" {
		fmt.Println("baron")
	} else {
		fmt.Println("tidak sama")
	}

	//var length = len(name)

	/*
		/ Short statement
	*/
	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
