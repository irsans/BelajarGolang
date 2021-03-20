package main

import "fmt"

//sebuah template data untuk menggabungkan 0 atau lebih tipe data lain nya dalam satu kesatuan
//data di struct disimpan dalam field
//struct kumpulan dari field
//sama seperti class
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var irsan Customer
	irsan.Name = "Taufik"
	irsan.Address = "Indonesia"
	irsan.Age = 30

	taufik := Customer{
		Name:    "Irsan",
		Address: "Jakarta",
		Age:     20,
	}

	fmt.Println(irsan)
	fmt.Println(taufik)
}
