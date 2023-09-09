package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("rezky")
	u := uuid.New()
	fmt.Println(u.String())

	fmt.Println(len("rezky"))
	fmt.Println("rezky"[0]) //result byte code
	fmt.Println("rezky"[1]) //result byte code
}
