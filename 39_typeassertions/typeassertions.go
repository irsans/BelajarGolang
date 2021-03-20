package main

import "fmt"

//type assertion adalah mengcover return data dari interface kosong

func random() interface{} {
	return "OK"
}
func main() {
	result := random()
	//type assertion di sesuaikan tipe data return nya agar tidak salah lebih baik meggunakan switch case
	resultString := result.(string)
	fmt.Println(resultString)
	//pastikan hasil type assesrtion nya sama dengan value tipe data return nya

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, " is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("unknown type")
	}
}
