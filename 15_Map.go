package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "irsan",
		"address": "karawaci",
	}
	person["title"] = "Programmer"

	delete(person, "title")
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar GoLang"
	book["author"] = "irsan"
	book["ups"] = "salah"

	book["ups"] = "benar"

	fmt.Println(book)
}
