package main

import "fmt"

// นิยามค่าคงที่
func main() {
	var name string = "fah"
	age := 21
	var gender string = "M"
	fmt.Printf("Type name is %T\n", name)
	fmt.Printf("Value name is %v\n", name)
	fmt.Println("")
	fmt.Printf("Type age is %T\n", age)
	fmt.Printf("Value age is %v\n", age)
	fmt.Println("")
	fmt.Printf("Type gender is %T\n", gender)
	fmt.Printf("Value gender is %v\n", gender)
}
