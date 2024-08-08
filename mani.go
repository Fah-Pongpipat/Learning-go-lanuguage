package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

func main() {
	var person1 Person
	person1.name = "Fah"
	person1.age = 21
	person1.gender = "M"
	fmt.Println(person1)

	person2 := Person{name: "Ed", age: 12, gender: "M"}
	fmt.Println(person2)
}
