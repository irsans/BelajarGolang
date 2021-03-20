package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perlungan ke", counter)
		counter++
	}

	//didalam for ada 2 statement init statement(sebelum kondisi, di ekse 1x)  dan post statement(selalu di eksekusi di akhir tiap looping, diexe sebanyak looping nya)

	//init statement

	/*for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke berapa", counter)
	}*/

	slice := []string{"Taufik", "Irsan", "Eko", "Patrio"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])

	}

	for index, value := range slice {
		fmt.Println("Index", index, "=", value)

	}
	//_ kalau index dari array nya tidak digunakan dan hanya ingin mengambil value nya saja karena golang jika ada 1 variabel yang tidak dipakai dia otomatis error
	for _, value := range slice {
		fmt.Println(value)

	}

	person := make(map[string]string)
	person["name"] = "Irsan"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
