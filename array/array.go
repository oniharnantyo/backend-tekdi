package main

import "fmt"

func main()  {
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Watermelon"

	fmt.Println(fruits)

	var vehicles []string
	vehicles = append(vehicles, "Car")
	vehicles = append(vehicles, "Motorcycle")
	vehicles = append(vehicles, "Bus")

	for i, vehicle := range vehicles {
		fmt.Println(i, vehicle)
	}

	var vehiclesAnother = make([]string,10)
	vehiclesAnother[0] = "Car"

	fruitsAnother := []string{"Apple", "Banana", "Watermelon"}

	fmt.Println("Result", fruitsAnother[2:3])

	var numbers []int
	fmt.Println(len(numbers))
	for i := 0 ; i < 10 ; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
}