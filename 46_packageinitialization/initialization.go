package main

import (
	_ "LearningGolang/database"
	_ "fmt"
)

//jika kita mengimport sebuah package namun tidak di gunakan kita bisa menambahkan tanda _ underscore di depan nama package import nya
// _ namanya blank identifier , kalau tidak pake tanda _ di hapus

func main() {
	// result := database.GetDatabase()
	// fmt.Println(result)
}
