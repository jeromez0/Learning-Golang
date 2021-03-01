package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter a string: ")
	
	str, err := reader.ReadString('\n')
	if err != nil {
		panic("error: newline not found")
	}
	fmt.Print(str)
}
