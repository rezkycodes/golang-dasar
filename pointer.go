package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountry(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	addr1 := Address{"Gorontalo", "Gorontalo", "Indo"}
	addr2 := &addr1

	addr2.City = "Bandung"

	//addr2 = &Address{"Malang", "Jawa Timur", "Indo"}
	*addr2 = Address{"Malang", "Jawa Timur", "Indo"}

	fmt.Println(addr1)
	fmt.Println(addr2)

	var address4 *Address = new(Address)
	address4.City = "Boalemo"
	fmt.Println(address4)

	var alamat = Address{
		City:     "bone bolango",
		Province: "gorontalo",
		Country:  "",
	}
	//ChangeCountry(&alamat)
	var alamatPointer *Address = &alamat
	ChangeCountry(alamatPointer)
	fmt.Println(alamatPointer)
}
