package main

import "fmt"

func main() {
	Greet()             // вызов функции
	PersonGreet("Oleg") //передача строки/аргумента на вход функции
	//CTRL+ кликнуть по функции - переносит к функции
	FioGreet("Oleg", "Kovi")

	Return_sum := sum(1, 5)
	// можно присвоить через переменные:
	//first, second := 1,2
	//Return_sum :=sum(first,second)
	print(Return_sum)

	sum, multiply := SumAndMultiply(1, 2)
	fmt.Printf("\n SUM %v, \n multiply %v", sum, multiply)

	_, multiply64 := nameSumAndMultiply(1, 2)
	fmt.Println(multiply64)

}

/*
func funcName (params...) (returned values) {
	function`s body
}
*/

func Greet() {
	fmt.Println("Hello!!!")
}

func PersonGreet(name string) { // входной параметр названия name типа string
	fmt.Printf("Zdarova %s\n", name)

}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s \n", name, surname)
}

func sum(first, second int) int { //принимает 2 параметра типа int, возврвщает один типа int. Для возврата - return
	sum := first + second
	return sum
}

func SumAndMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func nameSumAndMultiply(first, second int) (sum, multiply int64) { // две переменных возвращается
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return // не обязательно уже писать переменные в return.
}
