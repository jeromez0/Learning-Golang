package main

import "fmt"

func main() {

	// a string is a slice of bytes:
	testString := "This is a test."
	
	fmt.Printf("The first word of a is \"%s\"\n", firstWord(testString))

}

func firstWord(str string) []byte {
	//word is an empty slice of bytes
	word := []byte{}
	
	//get everything in the argument up to but not 
	//including the first space
	for ch := range str {
		if str[ch] == ' ' {
			break
		} else {
			word = append(word, str[ch])
		}
	}
	return word
}
