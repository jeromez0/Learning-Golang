package main

import "fmt"

func main() {
	nextFibb := fibb()
	fmt.Print("1 1 ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", nextFibb())
	}
}

func fibb() func() uint {
	var prevf uint = 1
	var f uint = 1
	return func() uint {
		f, prevf = f + prevf, f
		return f
	}
}
