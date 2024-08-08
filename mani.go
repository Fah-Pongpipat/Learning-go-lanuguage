package main

import "fmt"

// call a function
func myMessage() {
	fmt.Println("I just got executed!")
}

// parameters
func myName(name string) {
	fmt.Println("Hello", name)
}

// function return and input parameters
func plus(a int, b int) int {
	return a + b
}

// function return
func printHello() string {
	return "Hello  world"
}
func main() {
	myMessage()
	myName("Fah")
	sum := plus(5, 6)
	fmt.Println(sum)
	fmt.Println(printHello())
}
