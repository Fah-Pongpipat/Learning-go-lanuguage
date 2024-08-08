package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 8 {
			fmt.Println("Ended run code")
			break
		} else {
			fmt.Println("Round :", i)
		}
	}
}
