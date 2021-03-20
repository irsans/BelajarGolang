package main

import "fmt"

func main() {
	name := "irsans"

	switch name {
	case "irsan":
		fmt.Println("Hello Irsan")
	case "taufik":
		fmt.Println("Hello taufik")
	default:
		fmt.Println("Hi")
	}

	//short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}
	//code switch tanda kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama lebih dari 5 karakter")
	case length < 5:
		fmt.Println("Nama di bawah 5 karakter")
	default:
		fmt.Println("Nama sudah benar")
	}
}
