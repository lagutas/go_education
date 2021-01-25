package main

import (
	"fmt"
	"math"
)

func main() {
	var areaOfCircle float64
	fmt.Print("Enter areo of circle: ")
	fmt.Scanln(&areaOfCircle)

	fmt.Printf("Diameter of circle is - %.2f, circumference is %.2f\n", math.Sqrt(areaOfCircle / math.Pi) * 2, 2 * math.Sqrt(areaOfCircle / math.Pi) * math.Pi)

}