package main

import "fmt"

func main() {
 	// Slices in Go
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]string, 3) // Create a slice of strings with length 3

	// Assign values to the string slice
	slice2[0] = "Hello"
	slice2[1] = "World"
	slice2[2] = "Go"

	// Print the slices
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	// Append to a slice
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Updated Slice1:", slice1)

	// Slice slicing
	subSlice := slice1[2:5] // Get a sub-slice from index 2 to 4
	fmt.Println("SubSlice:", subSlice)

}