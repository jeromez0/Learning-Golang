package main

import "fmt"

func main() {
	defer f1()
	f2()
}

func f1() {
	fmt.Println("Hello from f1")
}

func f2() {
	fmt.Println("Hello from f2")
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}
