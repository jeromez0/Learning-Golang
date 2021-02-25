package main

import "fmt"

func main() {
	x := 5

	fmt.Printf("%v! = %v\n", x, fact(x)); 
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
