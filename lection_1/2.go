package main

import "fmt"

var name string // переменная видна всем функциям в этом пакете

func main() {
	//fmt.Println("This is IT")
	var hello string
	var boolean_1 bool
	//короткий синтаксис объявления переменной и присваивания значений только внутри функций
	// bool := false

	var name string //переменная видна только внутри функции.

	fmt.Println(hello) // функция выводит строку в консоль. Выведет в консоль пустую строку(значение не присвоено)
	hello = "hello"

	fmt.Printf("Type:%T Value: %v\n", hello, hello) //printF - форматированный вывод
	//%T - вывести тип переменной, %v - вывести значение переменной
	// \n - переход на след строку
	fmt.Printf("Type:%T Value: %v\n", boolean_1, boolean_1) // стандартно значение boolean - false

	name = "Строка"
	fmt.Println(name)

}
