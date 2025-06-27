package main

import ("fmt" 
	"unsafe")

func main() {

	// pointers in Go
	var a int = 42
	var ptr *int = &a // ptr is a pointer to a
	fmt.Println("Value of a:", a)         // prints 42
	fmt.Println("Address of a:", &a)      // prints the memory address of a
	fmt.Println("Value of ptr:", ptr)     // prints the address of a
	fmt.Println("Value pointed by ptr:", *ptr) // prints 42, the value at the address stored in ptr

	*ptr = 100 // changing the value at the address pointed by ptr
	fmt.Println("New value of a:", a) // prints 100, as the value at the address has changed
	fmt.Println("New value pointed by ptr:", *ptr) // prints 100, as the value at the address has changed
	fmt.Println("Address of ptr:", &ptr) // prints the address of the pointer itself
	fmt.Println("Value of ptr after change:", ptr) // prints the same address, as ptr still points to the same location
	fmt.Println("Address of a after change:", &a) // prints the same address, as a's address hasn't changed

	fmt.Println("Size of int:", unsafe.Sizeof(a)) // prints the size of an int in bytes
	fmt.Println("Size of pointer:", unsafe.Sizeof(ptr)) // prints the size of a pointer in bytes
}