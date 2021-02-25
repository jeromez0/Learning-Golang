package main

import "fmt"

func main() {

	grades := [7]int{93, 72, 68, 54, 100, 43, 86}
	strScore := ""
	for score := range grades{
		switch {
			case grades[score] > 90:
				strScore = "A"	
			case grades[score] > 80:
				strScore = "B"
			case grades[score] > 70:
				strScore = "C"
			case grades[score] > 60:
				strScore = "D"
			case grades[score] > 50:
				strScore = "F"		
		}
		fmt.Printf("Socre: %v; Letter Grade: %v\n", grades[score], strScore)
	}
}
