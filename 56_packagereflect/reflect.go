package main

import (
	"fmt"
	"reflect"
)

func isValid(data interface{}) string {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return "Data field " + t.Field(i).Name + " kosong"
			}
		}
	}
	return "Data Terisi semua"
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
	Alamat      string `required:"true"`
}

func main() {
	sample := Sample{"Irsan"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	//sample.Name = ""
	fmt.Println(isValid(sample))

	contoh := ContohLagi{"taufik", "", "asd"}
	fmt.Println("ini hasil contoh", isValid(contoh))
	fmt.Println(reflect.TypeOf(contoh).NumField())
	value := reflect.ValueOf(sample).Field(0).Interface()
	fmt.Println("ini reflect value", value)
}
