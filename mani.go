package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter your number :")
	fmt.Scanln(&number)
	if number >= 1 {
		fmt.Println("มากกว่า 1 ")
	} else {
		fmt.Println("น้อยกว่า 1")
	}
}
