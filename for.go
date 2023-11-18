package main

import "fmt"

func main() {
	//counter := 1
	
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke : ", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke:", counter)
	}

	slice := []string{"Rezky", "Pradana", "Budihartono"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range with slices
	// tanda '_' jika variable tidak di gunakan
	//for _, value := range slice {
	//	//fmt.Println("index", i, "=", name)
	//	fmt.Println(value)
	//}

	person := make(map[string]string)
	person["name"] = "Rezky"
	person["address"] = "Gorontalo"
	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}
}
