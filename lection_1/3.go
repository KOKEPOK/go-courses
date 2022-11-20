package main

import "fmt"

const name = "Oleg" //Константа не меняет значение.

func main() {
	fmt.Println(name)

	hello := fmt.Sprintf("hello %s", name) //функция возвращающая строку
	_ = hello                              // лайфхак для неиспользуемых переменных. Перекладывание значение переменной пустому идентификатору
	// %s - для вывода строки
	fmt.Printf(hello)
	//fmt.Printf("Type %T Value: %v\n", name, name)

}
