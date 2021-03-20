package main

import "fmt"

func main() {
	var nilaiAkhir = 80
	var absensi = 80

	var LulusUjian = nilaiAkhir >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = LulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
