package main

import "fmt"

func main() {
	
	// a is a pointer to an int
	a := new(int)
	
	// assign 9 as the "contents of" a
	*a = 9
	
	//print the address and contents of a
	fmt.Printf("a is %v, *a is %v\n", a, *a)
	
	// b points to the same location as a
	b := a

	// print the address and contents of b
	fmt.Printf("b is %v, *b is %v\n", b, *b)
	
	// increment the contents of a by one 
	*a++
	
	// because b references the same memory as a,
	// changing a also changes b:
	fmt.Printf("b is %v, *b is %v\n", b, *b)

}
