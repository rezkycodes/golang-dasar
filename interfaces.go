package main

import "fmt"

// Kontrak awal
type HasName interface {
	GetName() string
	cekUmur() int
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName(), " | Umur  ", hasName.cekUmur())
}

type Person struct {
	Name string
	Umur int
}

// function person di bawah ini akan mengikuti kontrak dari interface di atas
func (person Person) GetName() string {
	return person.Name
}
func (person Person) cekUmur() int {
	return person.Umur
}

type Animal struct {
	Name string
	Umur int
}

func (animal Animal) GetName() string {
	return animal.Name
}
func (animal Animal) cekUmur() int {
	return animal.Umur
}
func main() {
	var rezky Person
	rezky.Name = "Rezky"
	rezky.Umur = 31

	SayHello(rezky)

	cat := Animal{
		Name: "Push",
		Umur: 1,
	}

	SayHello(cat)
}
