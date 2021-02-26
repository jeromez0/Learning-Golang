package main

import "fmt"

type Rect struct {
	w, h float64
}

func main() {
	// r1 is of type Rect
	var r1 Rect
	r1.w = 5.3
	r1.h = 2.7
	
	fmt.Printf("Area of r1 = %.2f\n", area(r1))
	
	// r2 is a pointer to a React
	r2 := new(Rect)
	r2.w = 3.5
	r2.h = 7.1
	
	fmt.Printf("Area of r2 = %.2f\n", area(*r2))

}

func area(r Rect) float64 {
	return r.w * r.h
}
