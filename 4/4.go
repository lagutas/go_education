package main

import (
	"fmt"
	"log"
)

//Познакомьтесь с алгоритмом сортировки вставками.
//Напишите приложение, которое принимает на вход набор целых
//чисел и отдаёт на выходе его же в отсортированном виде.
//Сохраните код, он понадобится нам в дальнейших уроках.

func main() {
	fmt.Print("Введите количество цифр в массиве: ")
	var numOfDigits uint32
	_, err := fmt.Scanln(&numOfDigits)
	if err != nil {
		log.Fatalf("Получили ошибку воода: %v", err)
	} else if numOfDigits == 0 {
		log.Fatalf("Введите больше 0 элементов")
	}

	sliceOfDigits := []int32{}

	fmt.Println("Введите элементы массива: ")
	var i uint32
	for i = 0;i <= numOfDigits-1;i++ {
		var digit int32
		_, err = fmt.Scanln(&digit)
		if err != nil {
			log.Fatalf("Получили ошибку воода: %v", err)
		}
		sliceOfDigits = append(sliceOfDigits, digit)
	}

	for i, _ := range(sliceOfDigits) {
		for j := i; j>0 && sliceOfDigits[j-1] > sliceOfDigits[j]; j-- {
			sliceOfDigits[j-1], sliceOfDigits[j] = sliceOfDigits[j], sliceOfDigits[j-1]
		}
	}

	fmt.Println("отсортированные элементы")
	for _, el := range(sliceOfDigits) {
		fmt.Println(el)
	}
}
