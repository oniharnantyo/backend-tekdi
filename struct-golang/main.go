package main

import (
	"fmt"
	"time"
)

// Definisi / blueprint
type Person struct {
	Name string
	DateOfBirth time.Time
	Gender string
	Hobby []string
	School School
}

func (p *Person)Print()  {
	fmt.Println(p)
}

func (p *Person)getName() string {
	return p.Name
}

func (p Person)updateName(name string)  {
	p.Name = name
}

func (p *Person)updateNameWithPointer(name string)  {
	p.Name = name
}

type School struct {
	Name string
	Address string
}

func main()  {
	person := Person{
		Name:        "Bob",
		DateOfBirth: time.Date(1990,1, 2, 0,0,0,0,time.Now().Location()),
		Gender:      "M",
		Hobby:       []string{"Football", "Singing"},
		School: School{
			Name:    "TK Maju",
			Address: "Jalan jalan",
		},
	}

	//&{Bob 1990-01-02 00:00:00 +0700 WIB M [Football Singing]}
	person.Print()

	//Bob
	fmt.Println("Name before update", person.Name)

	person.updateName("Martin")

	fmt.Println("Name after update :",person.Name)

	person.updateNameWithPointer("Martin")

	fmt.Println("Name after update with pointer:",person.Name)

	var persons []Person

	persons = append(persons, person)
	persons = append(persons, Person{
		Name:        "martin",
		DateOfBirth: time.Now(),
		Gender:      "M",
		Hobby:       []string{"basketball"},
	})

	fmt.Println(persons)

	fmt.Println("Bob school name:", person.School.Name)
}


kodePos  := map[int]string

type kodePos struct {
	kode int
	desa []string
}

kodePosList := []kodePos

15112 :> "Depok", "sewon"
15113 :> "Sewon", "ngaglik"
15114
15115
15116

kode[15112] = "Depok", "sewon"

x : 15116

for kodePoslist {
	if kodePos.kode == x {

	}
}