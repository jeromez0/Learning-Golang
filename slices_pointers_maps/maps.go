package main

import "fmt"

func main() {

	/* 
		a map is a table of key/value pairs. Maps
		are also known as hash tables, dictionaries, or 
		associative arrays in other languages.
	*/

	// syntax for declaring a map with make
	// digits := make(map[string]int)

	digits := map[string]int {
		"zero"  : 0,
		"one"   : 1,
		"two"   : 2,
		"three" : 3,
		"four"  : 4,		
	}
	digits["five"] = 5	// adding a new k/v pair
	
	if digits["four"] != 0 { // checking on the existence of a key
		fmt.Printf("that digit is %v\n", digits["four"])
	} else {
		fmt.Printf("that digit is not in the digits map.")
	}

	// this is called the comma ok idiom. It tests for the 
	// presence of a pair in the map
	
	key := "six"
	
	var digit int
	var ok bool
	
	if digit, ok = digits[key]; ok { // note syntax here
		fmt.Printf("%v is %v\n", key, digit)
	} else {
		fmt.Printf("%v is not a key in the map.\n", key)
	}
	// we can also throw away the value if we're only interested 
	// in knowing if the key exists.  this is done to avoid
	// the problem of initializing but not using a variable.
	
	if _, ok = digits["nine"]; ok {
		fmt.Println("nine is in the map.")
	} else {
		fmt.Println("nine is not in the map.")
	}
	
	// deleting a pair from a map:
	delete(digits, "zero")
	
	// we can print the entire map: 
	fmt.Printf("%v\n", digits)
	
		
}
