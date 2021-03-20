package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func getName(name string) string {
	return name
}
func main() {

	//function dapat di assign kedalam sebuah variabel lalu setelah di assign untuk memanggil func tersebut bisa dengan nama variabel yang telah di assign
	goodbye := getGoodBye
	result := goodbye("irsan")
	getname := getName(result)
	fmt.Println(goodbye("irsan"))
	fmt.Println(result)
	fmt.Println(getname)
}
