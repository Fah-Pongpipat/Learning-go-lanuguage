package main

import "fmt"

func main() {
	list := []int{10, 20, 30, 40, 50}

	for index, value := range list {
		fmt.Println("Index :", index, " Value :", value)
	}
	fmt.Println("---------------------------")
	for i := 0; i <= 10; i++ {
		if i == 8 {
			fmt.Println("Ended run code")
			break
		} else {
			fmt.Println("Round :", i)
		}
	}
}
