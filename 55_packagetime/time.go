package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	//format layout standar time.RFC3339
	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}
