package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortArrayByPoppingSmallest(arr []int) []int {
	// Create an empty slice to store the sorted elements
	var sorted []int

	// Outer loop continues while the input array is not empty
	for len(arr) > 0 {
		// Assume the first element is the smallest
		smallestIndex := 0
		smallest := arr[0]

		// Inner loop to find the smallest element
		for i := 1; i < len(arr); i++ {
			if arr[i] < smallest {
				smallest = arr[i]
				smallestIndex = i
			}
		}

		// Append the smallest element to the sorted list
		sorted = append(sorted, smallest)

		// Remove (pop) the smallest element from the original array
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}

	// Return the sorted array
	return sorted
}

// Function to generate a random array with n elements
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(1000000) // Random number between 0 and 999999
	}
	return arr
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create arrays with 10,000, 100,000, and 1 million elements
	arr1 := generateRandomArray(10000)
	arr2 := generateRandomArray(100000)
	arr3 := generateRandomArray(1000000)

	// Measure and print the time for sorting arr1 (10,000 elements)
	start := time.Now()
	sortArrayByPoppingSmallest(arr1)
	elapsed := time.Since(start)
	fmt.Printf("Time taken to sort 10,000 elements: %s\n", elapsed)

	// Measure and print the time for sorting arr2 (100,000 elements)
	start = time.Now()
	sortArrayByPoppingSmallest(arr2)
	elapsed = time.Since(start)
	fmt.Printf("Time taken to sort 100,000 elements: %s\n", elapsed)

	// Measure and print the time for sorting arr3 (1 million elements)
	start = time.Now()
	sortArrayByPoppingSmallest(arr3)
	elapsed = time.Since(start)
	fmt.Printf("Time taken to sort 1 million elements: %s\n", elapsed)
}
