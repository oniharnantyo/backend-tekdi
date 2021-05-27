package main

import "fmt"

func main()  {
	// Numerik non desimal
	var angka int32 = -1
	var positif uint = 10

	// Numerik desimal
	var desimal float64 = 10.1

	// Boolean
	var boolean = true

	// String / Huruf
	var huruf string  = "hello"

	fmt.Println(fmt.Sprintf(`Numerik: %d`, angka))
	fmt.Println(fmt.Sprintf(`Numerik positif: %d`, positif))
	fmt.Println(fmt.Sprintf(`Desimal: %f`, desimal))
	fmt.Println(fmt.Sprintf(`Bool: %t`, boolean))
	fmt.Println(fmt.Sprintf(`Huruf: %s`, huruf))
}