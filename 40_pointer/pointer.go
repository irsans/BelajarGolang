package main

import "fmt"

type address struct {
	//
	City, Province, Country string
}

func main() {
	address1 := address{
		City:     "Tangerang",
		Province: "Banten",
		Country:  "Indonesia",
	}
	//pointer harus menggunakan tanda & agar variable pertama juga berubah
	address2 := &address1
	address3 := &address1
	address2.City = "Bandung"

	//tanda * ini  di gunakan ketika kita ingin mengubah data dari semua field address2 yang mengarah ke pointer address 1 sehingga ketika data di ubah
	//address 1 pun akan ikut berubah mengacu ke reference yang sama. karena sebelum nya dia mengarah ke data struct yang pertama(address1)
	*address2 = address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}
	//jika ingin membuat data baru maka pada = Address cukup di tambahkan tanda & agar menjadi data memory yang baru
	address4 := &address{
		City:     "Lamongan",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	//membuat pointer dengan data kosong menggunakan new

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	address5 := new(address)
	address5.City = "Jakarta"
	fmt.Println(address5)

}
