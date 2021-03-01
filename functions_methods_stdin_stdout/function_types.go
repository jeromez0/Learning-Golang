package main

import "fmt"

func main() {

	var a, b int
	add := func() int {
		return a + b
	}
	
	a, b = 3, 9
	fmt.Printf("%v + %v = %v\n", a, b, add())
	a, b = 4, -7
	fmt.Printf("%v + %v = %v\n", a, b, add())
}
