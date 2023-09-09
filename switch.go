package main

import "fmt"

func main() {
	var name = "Rezkysssss"

	switch name {
	case "baron":
		fmt.Println("baron")
	case "Rezky":
		fmt.Println("rezky")
	default:
		fmt.Println("no name")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	//switch tanpa kondisi
	length := len(name)

	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
