package main

import "fmt"

func main() {
	var digit uint64

	fmt.Print("Enter digit: ")
	fmt.Scanln(&digit)

	fmt.Printf("Num of hundred - %d, num of dozens - %d, num of units - %d\n", digit / 100, (digit - (digit / 100) * 100) / 10, digit - (digit / 100) * 100 - (digit - (digit / 100) * 100) / 10 * 10)
}
