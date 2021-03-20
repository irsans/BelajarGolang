package main

import (
	"fmt"
	"strconv"
)

func main() {
	//hampir semua function di stringconv bisa mengembalikan value error

	boolean, err := strconv.ParseBool("salah")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) // mengubah string ke int64
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) //mengubah int64 ke string
	fmt.Println(value)

	value1 := strconv.Itoa(1000)
	fmt.Println(value1)
}
