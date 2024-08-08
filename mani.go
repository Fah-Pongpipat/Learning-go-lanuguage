package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter your score :")
	fmt.Scanln(&score)
	switch score {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("Error")
	}
}
