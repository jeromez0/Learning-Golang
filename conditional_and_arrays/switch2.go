package main

import "fmt"

func main() {

	//loop from 1 to 10
	for i:=1; i <= 10; i++{
	// assign a numSize to i
		switch i {
			case 1,2,3:
				fmt.Printf("%v is a small number.\n",i)
			case 4,5,6,7:
				fmt.Printf("%v is a medium number.\n",i)
			default:
				fmt.Printf("%v is a large number.\n",i)
		}
	}
}
