package main

import "fmt"

func main() {
	i := 0
	i += 1 // тоже самое что и i = i + i
	i++    //инкремент. увеличение значения на единицу.
	fmt.Println(i)
	i = 10
	i = i - 1
	i -= 1 // i= i-1
	i--    // декремент
	fmt.Println(i)

	// циклы
	//for init; condition; post statement. init - начальное значение. condition - условие. post statement - После выполнения что сделать
	var j int                // Аналогично:
	for j = 0; j < 10; j++ { // for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	sum := 1
	for sum < 20 { // init и post statement - не обязательное
		sum += sum
		fmt.Println(sum)
	} // в этом цикле накопится sum до 32.

	//for как while:
	for sum <= 100 { //32 <=100
		sum += 1
		fmt.Println(sum)
	} // выведет значения от 32 до 101.

	//for { // бесконечный цикл
	//	fmt.Println("Stop me please")
	//}
}
