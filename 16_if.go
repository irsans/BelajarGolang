package main

import "fmt"

func main() {
	name := "taufik"

	if name == "irsan" {
		fmt.Println("ini true")
	} else if name == "taufik" {
		fmt.Println("ini true 2")
	} else {
		fmt.Println("ini false")
	}

	var length = len(name)

	if length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
