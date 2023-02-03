package main

import "fmt"

func main() {
	array := [5]string{"I", "am", "stupid", "and", "weak"}

	for i, v := range array {
		if v == "stupid" {
			array[i] = "smart"
		}

		if v == "weak" {
			array[i] = "strong"
		}
	}

	fmt.Println(array)
}
