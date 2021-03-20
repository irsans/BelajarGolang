package main

import "fmt"

//nill hanya bisa di gunakan di beberapa tipe data seperti interface, function,map, slice

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	person := newMap("")
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}

}
