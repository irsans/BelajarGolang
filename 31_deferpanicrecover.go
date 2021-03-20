package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	//ketika terjadi error yang di bawah nya tidak tereksekusi , jika ingin tetap memanggil dengan menggunakan defer function
	fmt.Println("Result", result)
	//logging()
}

func endAppPanic() {

	//recover menangkap data panic ,dengan recover proses panic akan terhenti sehingga program tetap akan berjalan
	//pastika recover di ekseksu di end app agar bisa di gunakan
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func rundAppPanic(error bool) {
	defer endAppPanic()
	if error {
		panic("APLIKASI ERRROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApplication(10)
	rundAppPanic(true)
	fmt.Println("Irsan")
}
