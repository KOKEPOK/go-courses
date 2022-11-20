package main

import "fmt"

func main() {
	age := 15

	if age < 18 {
		fmt.Println("You are to young (full)")
	}
	if isChild := isChildren(age); isChild == true { // переменная внутри if доступна только в if
		fmt.Println("You are very young(short)")
	}

	age = 20
	if age < 18 {
		fmt.Println("You are to young")
	} else {
		fmt.Println("You are an adult")
	}
	if age >= 6 && age <= 18 {
		fmt.Println("You are in school")
	}
	//
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You have to get new passport")
	}

	if !isChildren(age) {
		fmt.Println("You are an adult (ischildren false)")
	}

}
func isChildren(age int) bool {
	return age < 18
}
