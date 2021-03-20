package main

import "fmt"

//tipe data return nya string
func getHello(name string) string {
	if name == "" {
		return "Hello" + name
	} else {
		return "Hello Bro"
	}
}

func main() {
	result := getHello("Taufik")
	fmt.Println(result)
}
