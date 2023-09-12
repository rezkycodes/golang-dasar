package main

/**
// tambahkan blank identifier jika package akan di
// panggil tanpa memangil method GetDatabase karena defaultnya
// golang akan eror jika ada package yg tdk di gunakan
/ */
import (
	_ "belajar-golang/database"
	// "fmt"
)

func main() {
	// result := database.GetDatabase()
	// fmt.Println(result)

}