package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ { //5 %2 = 1

		if i%2 == 1 { // получить остаток от деления. Если остатка нет , то =0.
			continue // прерываем данную операцию, дальше не идём
		}
		fmt.Println(i)
	}

Label1:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I: ", i, "J:", j)
			if i >= 10 {
				continue Label1
			}
		}
	}

}
