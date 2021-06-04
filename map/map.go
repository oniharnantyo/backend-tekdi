package main

import "fmt"

func main() {
	var person map[string]string

	person = map[string]string{
		"name":    "Oni",
		"address": "Babarsari",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var person1 map[string]interface{}
	person1 = map[string]interface{}{
		"name": "Halo",
		"age":  1,
	}

	var person2 map[string]interface{}
	person2 = map[string]interface{}{
		"name": "Bob",
		"age":  1,
	}

	var persons []map[string]interface{}
	persons = append(persons, person1)
	persons = append(persons, person2)

	fmt.Println("Name", persons[1]["name"])
	fmt.Println("age", persons[1]["age"])

	if fmt.Sprintf(`%v`, persons[0]["name"]) == "bob" {

	}
	/*
		variable x dan y
		nanti akan ada perintah:
			- sum : tambah
			- mul : dikali

		akan ada banyak variable x dan y

		sum : tambah
		mul : kali

		data :
		{
			x : 10, y = 5, operation = sum
			x : 5, y = 3, operation = mul
			x : 3, y = 2, operation = sum
			x : 6, y = 9, operation = mul
		}

		output :
			Hasil perhitungan
	*/
}
