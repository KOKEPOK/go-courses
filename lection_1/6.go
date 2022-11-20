package main

import "fmt"

func main() {
	first, second := 1, 2

	var multiplier func(x, y int) int                // переменная с типом функции
	multiplier = func(x, y int) int { return x * y } // а потом присвоить ей функцию))
	fmt.Println(multiplier(first, second))

	divider := func(x, y int) int { return x / y }
	fmt.Println(divider(second, first)) // на вход переменной подаем 2 числа: 2, 1.
	//возвращаемое значение: 2/1 = 2.

	sumFunc := func(x, y int) int {
		return x + y
	}
	SubtractFunc := func(x, y int) int {
		return x - y
	}
	fmt.Println(calculate(first, second, sumFunc))
	fmt.Println(calculate(second, first, SubtractFunc))

	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)
	fmt.Println(divideBy2(100))
	fmt.Println(divideBy10(100))
}

func calculate(x, y int, action func(x, y int) int) int { //Функция может принимать функцию:
	return action(x, y) // action это жействие.
}

// возврат функции из функции
func createDivider(divider int) func(x int) int { // divider = 2. x = возвращаемый параметр который входной.
	dividerFunc := func(x int) int { //или возврат функции из функции
		return x / divider
	}
	return dividerFunc
}
