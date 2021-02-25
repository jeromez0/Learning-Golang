package main

import "fmt"

func main() {
	a, b := 2, 3
	
	if a == b {
		fmt.Printf("%v and %v are equal\n",a,b)
	}else if a > b{
		fmt.Printf("%v is greater than %v\n",a,b)
	}else {
		fmt.Printf("%v is less than %v\n",a,b)
	}
}
