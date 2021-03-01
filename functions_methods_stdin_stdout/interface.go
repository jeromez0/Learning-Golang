package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Cylinder struct {
	radius float64
	height float64
}

type Describer interface {
	describe()
}

func (c Circle) describe() {
	area := math.Pi * math.Pow(c.radius, 2)
	circ := 2 * math.Pi * c.radius
	fmt.Printf("A circle with radius %.2f.\narea %.2f.\nand circumference %.2f\n\n", c.radius, area, circ)
}

func (c Cylinder) describe() {
	volume := math.Pi * math.Pow(c.radius, 2) * c.height
	circ := 2 * math.Pi * c.radius
	fmt.Printf("A cylinder with radius %.2f,\nvolume %.2f,\n and circumference %.2f\n\n", c.radius, volume, circ)
}

func main() {
	circle := Circle{5}
	cylinder := Cylinder {5, 3}
	circle.describe()
	cylinder.describe()
}

