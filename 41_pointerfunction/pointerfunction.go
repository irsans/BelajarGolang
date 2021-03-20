package main

import "fmt"

type Address struct {
	City, Province, Country string
}

//agar data yang di kirim dapat berubah di pointer nya di tambahkan *
func changeAddressToIndonesia(address *Address) {
	address.Country = "Jepang"
}

func changeAddressToIJepang(address1 *Address) {
	address1.Country = "Jepang1"
}

func main() {
	address := &Address{"Tokyo", "TokyoMetro", ""}
	address1 := &Address{"Tokyo1", "TokyoMetro1", ""}
	changeAddressToIndonesia(address)
	changeAddressToIJepang(address1)
	fmt.Println(address) // tidak berubah
	fmt.Println(address1)
}
