package main

import "fmt"

func main()  {
	task1 := 80
	task2 := 50
	task3 := 30

	taskAverage := (task1 + task2 +task3 ) / 3

	fmt.Println(fmt.Sprintf("Average: %d", taskAverage))
	if  taskAverage > 75{
		fmt.Println("Graduate")
	} else if taskAverage > 50{
		fmt.Println("Remidial")
	} else {
		fmt.Println("Failed")
	}

	point := 5

	switch point {
	case 5:
		fmt.Println("Bad")
	case 6,7,8:
		fmt.Println("Not bad")
	case 9:
		fmt.Println("Good")
	}
}
