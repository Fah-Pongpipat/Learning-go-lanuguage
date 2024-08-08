package main

import "fmt"

// Variadic function
func variadic_func(number ...int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}

func main() {
	result := variadic_func(10, 20, 30)
	fmt.Println("Sum :", result)
}
