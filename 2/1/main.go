package main

import "fmt"

func main() {
	var a float32
	var b float32

	fmt.Print("Enter a side: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b side: ")
	fmt.Scanln(&b)

	fmt.Printf("Area of rectangle is %.2f\n", a * b)
}