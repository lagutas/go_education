package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b, res float32
	var op string


	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalf("Получили ошибку воода: %v", err)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.Fatalf("Получили ошибку воода: %v", err)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			log.Fatal("Делить на 0 нельзя")
		} else {
			res = a / b
		}
	default:
		log.Fatal("Операция выбрана неверно")
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)



}
