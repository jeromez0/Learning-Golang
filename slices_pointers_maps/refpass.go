package main

import "fmt"

func main() {
	num := 5
	negate(&num)
	fmt.Printf("num is now %v\n", num)
}

func negate(a *int) {
	*a *= -1
}
