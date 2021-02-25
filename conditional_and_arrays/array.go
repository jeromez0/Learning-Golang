package main

import "fmt"

func main() {
	//initialize an array of test scores:
	scores := [8]float64{87.3, 92.0, 100, 78.9, 84, 98, 72.6, 64}
	
	//print the array
	fmt.Printf("Scores: %v\n", scores)

	//calculate the average:
	var avg float64 = 0.0 //inferred to be float64

	// add the scores
	for i:= range scores {
		avg += scores[i]
	}
	
	//divide the sum by 8;
	avg = avg / float64(len(scores))

	// print the average
	fmt.Printf("Average score: %.2f\n", avg)
}
