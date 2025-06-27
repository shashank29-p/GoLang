package main

import "fmt"

func main() {
	// Maps in Go
	map1 := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	map2 := make(map[string]string) // Create a map with string keys and values

	// Assign values to the string map
	map2["name1"] = "John"
	map2["name2"] = "Doe"
	map2["name3"] = "Smith"
	// Print the maps

	fmt.Println("Map1:", map1)

	fmt.Println("Map2:", map2)

	// Accessing map elements

	fmt.Println("Alice's age:", map1["Alice"])
	fmt.Println("Name1:", map2["name1"])

	// Adding a new key-value pair to map1
	map1["David"] = 40
	fmt.Println("Updated Map1:", map1)

	// Deleting a key-value pair from map2
	delete(map2, "name2")
	fmt.Println("Updated Map2 after deletion:", map2)

	// Iterating over a map

	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	for key, value := range map2 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Checking if a key exists in a map
	if age, exists := map1["Alice"]; exists {
		fmt.Printf("Alice's age exists: %d\n", age)
	} else {
		fmt.Println("Alice's age does not exist")
	}
	
   	if name, exists := map2["name2"]; exists {
		fmt.Printf("Name2 exists: %s\n", name)
	} else {
		fmt.Println("Name2 does not exist")
	}
}