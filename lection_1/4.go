package main

import "fmt"
import "unsafe" // для вычисления длины строки

func main() {
	var num1 int = 15
	var num2 int64 = 15
	var num3 uint8 = 1
	var num4 uint64 = 1
	//var num2 int = 15 // INT64 сложить с типом INT нельзя. Необходимо приведение типов
	// num2 := 15

	//fmt.Println(num1+num2)
	fmt.Println(int64(num1) + num2) // num1 преобразовали к тип int64. num2 - тип уже int64.
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))
}
