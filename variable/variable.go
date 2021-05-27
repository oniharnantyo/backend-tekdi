package main

import (
	"fmt"
	"reflect"
)

func main()  {
	//Variabel
	var name string
	var height, weight int

	name = "John"
	height = 170
	weight = 60

	gender := 10.1

	fmt.Println("Type : ",reflect.TypeOf(gender))

	// Konstanta
	// Nilainya tidak dapat diubah
	const country = "Indonesia"


	fmt.Printf("Hi %s", name)
	fmt.Println()
	fmt.Printf("Gender : %s", gender)
	fmt.Println()
	fmt.Printf("Height : %d cm", height)
	fmt.Println()
	fmt.Printf("Weight : %d kg", weight)
	fmt.Println()
	fmt.Printf("Country : %s", country)
}