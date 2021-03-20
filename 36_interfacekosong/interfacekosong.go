package main

import "fmt"

//interface kosong return dapat mereturn semua tipe data dan dapat menerima parameter apapun
func ups(i int) interface{} {
	//
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}
func main() {
	kosong := ups(2)
	fmt.Println(kosong)
}
