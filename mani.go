package main

import "fmt"

func main() {
	myCountry := map[string]string{"Thai": "ไทย", "Japanese": "ญี่ปุ่น", "Italy": "อิตาลี"}
	// fmt.Println(myCountry["Thai"])

	find := "Th"
	value, check := myCountry[find]
	if check {
		fmt.Print(value)
	} else {
		fmt.Println("Data not found ")
	}
}
