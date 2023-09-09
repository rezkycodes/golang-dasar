package main

import "fmt"

func main() {

	//math - operator
	var a = 1
	var b = 2

	var c = a + b
	var d = 5 * 2
	fmt.Println(c)
	fmt.Println(d)

	//augmented assignment
	var i = 10
	i += 10

	fmt.Println(i)

	//unary operator
	i++
	fmt.Println(i)

	var negative = -100
	var positif = +100
	fmt.Println(negative)
	fmt.Println(positif)
}
