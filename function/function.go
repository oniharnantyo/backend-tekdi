package main

import (
	"fmt"
	"pertemuan-2/function/print"
)

func main()  {
	// call from same package
	printHelloWorld()

	// call from another package
	print.PrintHelloWorld()

	someText := "Testing param"
	printText(someText)

	fmt.Println(returnHelloWorld())

	x := 10
	y := 20
	fmt.Println("Result:", sum(x, y))

	result, isPositive := sumAndCheckIsPositive(x,y)
	fmt.Println(fmt.Sprintf(`Result : %d, isPositive %t`, result, isPositive))

	printText(returnHelloWorld())
}

func printHelloWorld()  {
	fmt.Println("Hello world")
}

func printText(text string) {
	fmt.Println(text)
}

func returnHelloWorld()string  {
	return "Hello from return"
}

func sum(x,y int)int  {
	result := x + y

	return result
}

func sumAndCheckIsPositive(x,y int)(int, bool)  {
	result := x + y

	var isPositive bool

	if result >= 0 {
		isPositive = true
	}else {
		isPositive = false
	}

	return result, isPositive
}

func sumIntAndFloat(x int, y float64)float64  {
	result := float64(x) + y

	return result
}
