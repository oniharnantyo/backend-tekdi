package main

import "fmt"

func main()  {
	var name string
	var height, weight int

	name = "John"
	height = 170
	weight = 60
	gender := "Male"


	fmt.Printf("Hi %s", name)
	fmt.Println()
	fmt.Printf("Gender : %s", gender)
	fmt.Println()
	fmt.Printf("Height : %d cm", height)
	fmt.Println()
	fmt.Printf("Weight : %d kg", weight)
}