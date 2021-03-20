package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(slice1)

	//months = append(months, "Tambah")
	//fmt.Println(slice1)

	//slice1[0] = "Mei Ubah"
	//fmt.Println(months)

	var slice2 = months[2:8]
	fmt.Println("ini slice2", slice2)

	var slice3 = append(slice2, "Irsan")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println("ini slice3", slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//membuat slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "irsan"
	newSlice[1] = "Taufik"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//mengcopy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	slice01 := days[2:]
	fmt.Println("ini silce01", slice01)

	slice02 := append(slice01, "weton")
	slice02[1] = "bukan minggu"
	fmt.Println("ini slice02", slice02)
}
