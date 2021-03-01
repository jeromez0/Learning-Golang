package main

import "fmt"


func main() {

	int1, int2 := 0, 0
	fmt.Printf("Enter two ints, separated by a space: ")
	
	count,err := fmt.Scan(&int1, &int2)
	
	if err != nil {
		panic(err)
	}
	fmt.Printf("You entered %v and %v. Total is %v. There are %v integers.\n", int1, int2, int1 + int2, count)

}
