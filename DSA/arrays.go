package main

import (
	"fmt"
)

func array() {
	var myArray = [5]int{0, 1, 2, 3, 4}
	fmt.Println("Array with default values", myArray)
	for index, value := range myArray {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

}
