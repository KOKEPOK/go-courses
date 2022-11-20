package main

import "fmt"

func main() {
	dollar := 30
	getDollarValue := func() int {
		return dollar // замыкает переменную. Просмотр актуального значения переменной
	}

	println(getDollarValue())
	dollar = 70
	fmt.Println(getDollarValue())
}
