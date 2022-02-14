package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age) + " years old"
}

func (p *Person) hasBirthDay() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "man" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "steven", lastName: "besselink",
		city: "oploo", gender: "man", age: 17}

	person2 := Person{firstName: "jose", lastName: "besselink",
		city: "oploo", gender: "vrouw", age: 47}
	// andere manier
	//person1 := Person{ "steven", "besselink",
	//	"oploo", "man", 17}

	//fmt.Println(person1)
	//fmt.Println(person1.age)
	person1.hasBirthDay()

	person2.getMarried("janssen")

	fmt.Println(person2.greet())
}
