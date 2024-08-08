package main

import "fmt"

func main() {
	my_Slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// before
	fmt.Println(my_Slice1)
	fmt.Println(len(my_Slice1))
	my_Slice1 = append(my_Slice1[:6], my_Slice1[6+1:]...)
	// after
	fmt.Println(my_Slice1)
}
