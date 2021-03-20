package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database house")
	var user *string = flag.String("user", "root", "Put your database house")
	var password *string = flag.String("password", "root", "Put your database house")

	//parse sangat penting agar data bisa di input
	flag.Parse()

	//karena ini pointer pada variable yang di panggil harus menambahkan tanda *
	//jika ada variable int lalu di inputkan string maka akan error dan muncul helper nya
	//cara memasukan value via command prompt go run flag.go -host=irsan -user=taufik -password=123 jika tidak di panggil seperti ini dia akan mengambil default value dari variabel nya
	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)
}
