package main

import "fmt"

func main() {
	fmt.Println("Введите три числа:")
	var first, second, third int

	fmt.Scan(&first)
	fmt.Scan(&second)
	fmt.Scan(&third)

	hasPositive := false
	if first > 0 || second > 0 || third > 0 {
		hasPositive = true
	}
  
	if hasPositive {
		fmt.Println("Одно из чисел — положительное")
	} else {
		fmt.Println("Среди чисел нет положительного")
	}

}
