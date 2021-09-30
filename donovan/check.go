package main

import(
	"fmt"
	"strconv"
)

// Читает число >= 0 и возвращает его
func CheckInputIntJNBE() int {
	var temp	string
	for {
		fmt.Scanf("%s\n",&temp)
		rez, err := strconv.Atoi(temp)
		if err == nil {
			if rez<1 {
				fmt.Println("! Число не может быть отрицательным или равно нулю.")
			} else {
				return rez
			}
		} else {
			fmt.Println("! Необходимо ввести целое положительное число.")
		}
	}
}

// Читает число >= 0 и возвращает его строкой
func CheckInputStringJNBE() string {
	return strconv.Itoa(CheckInputIntJNBE())
}