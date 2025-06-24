package main
import (
	"fmt" 
"strings"
)

	// Function with multiple return values
	func add(input1 int, input2 int) (int,int) {
		// can also write : add(input1, input2 int) 
		return input1 + input2,input1
	}

	// Function with named return values
	func sub(input1 int, input2 int) (result int) {
		return input1 - input2
	}

func main() {
	//Datatypes and Variables
	 var x string = "My name is john"
	 const constValue int =10
	 var floatValue float64 = 20.5
	 var boolVal bool=true
	 var intVersionAscii rune = 'A'


	 fmt.Printf("The value of x is: %s\n", x)
	 fmt.Printf("The type of x is: %T\n", x)
	 fmt.Printf("The value of constValue is: %d\n", constValue)
     fmt.Printf("The value of floatValue is: %f\n", floatValue)
	 fmt.Printf("The value of boolVal is: %t\n", boolVal)
	 fmt.Printf("The value of intVersionAscii is: %c\n", intVersionAscii)
    
	 // Conditional Statements
	 if  y:= 50; y <=40{
		fmt.Println("y is less than or equal to 40")
	 }else if y > 40 && y <= 60 {
		fmt.Println("y is between 40 and 60")
	 }else{
        fmt.Println("y is greater than 40")
	 }
	 
	 // Looping
	 for i := 0; i < 5; i++ {
		fmt.Printf("The value of i is: %d\n", i)
	 }

	// switch case
	switch intVersionAscii {
	case 'A':
		fmt.Println("The character is A")
	case 'B':
		fmt.Println("The character is B")
	default:
		fmt.Println("The character is neither A nor B")
	}


	// String Manipulation
	a := strings.Replace(x, "john", "doe", 1)
	b := x + " and I am a programmer"
	c := strings.Split(b, " ")
	

	fmt.Printf("The value of a, b and c is: %s, %s, %v\n", a, b, c)

    // Function calls
	resultAdd, _ := add(10, 5) // discard second return
	resultSub := sub(10, 5)
	fmt.Printf("The result of addition is: %d and subtraction is: %d\n", resultAdd, resultSub)
}