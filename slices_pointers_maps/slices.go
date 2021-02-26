package main

import "fmt"

func main() {
	// make a slice of int with
	// an initial length of 10 and 
	// a capacity of 100:
	
	a:= make([]int,10,100)
	
	i := 0
	for i < 10 {
		a[i] = 2 * i
		i++
	}
	fmt.Printf("a is %v\n", a)
	fmt.Printf("adding an element to a...\n")
	
	a = append(a, 20)
	
	fmt.Printf("a is now %v\n", a)
	b := a[1:5]
	fmt.Printf("b is %v\n", b)
	
	
}
