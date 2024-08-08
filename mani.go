package main

import "fmt"

func main() {
	list := []int{10, 20, 30, 40, 50}

	for _, value := range list {
		fmt.Println(" Value :", value)
	}

}
