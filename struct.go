package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct method biasa
//func  sayHi(customer Customer, name string) {
//	fmt.Println("hello", name, "My name is", customer.Name)
//}

// function struct
func (customer Customer) sayHi(name string) {
	fmt.Println("hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("huuuuu from", a.Name)
}
func main() {
	var rezky Customer

	rezky.Name = "Rezky P Budihartono"
	rezky.Address = "Gorontalo"
	rezky.Age = 30

	//panggil struct method biasa
	//sayHi(rezky, "keenan")

	//panggil struct method
	rezky.sayHi("keenan")
	rezky.sayHuu()

	//fmt.Println(rezky.Name)
	//fmt.Println(rezky.Address)
	//fmt.Println(rezky.Age)
	//
	////var baron Customer = Customer{}
	//baron := Customer{
	//	Name:    "Baron",
	//	Address: "JDS",
	//	Age:     0,
	//}
	//
	//fmt.Println(baron)
	//
	//budi := Customer{"Budi", "Gorontalo", 35}
	//fmt.Println(budi)
}
