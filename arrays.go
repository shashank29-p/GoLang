package main

import "fmt"

func main() {
   arr := []int{1, 2, 3, 4, 5}
   var names = [3]string{"Alice", "Bob", "Charlie"}
   cities := [...]string{"Delhi", "Mumbai", "Chennai"} //spread operator
   for i, v:= range arr {
	  fmt.Println("Index:", i, "Value:", v)
   }
   for i, name := range names {
	  fmt.Println("Index:", i, "Name:", name)
   }

   cities[0] = "Bangalore" // Modifying an array element

   fmt.Println("Cities:", cities)
}