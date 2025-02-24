package exercise04

import "fmt"

func Examples_conditionals() {

	// example if
	number := 10
	if number > 5 {
		fmt.Println("number %i is greater than 5", number)
	}
	// example if else
	if number > 15 {
		fmt.Println("number %i is greater than 15", number)
	} else {
		fmt.Println("number %i is less than 15", number)
	}

	// example switch
	switch number {
	case 1:
		fmt.Println("number is 1")
	case 2:
		fmt.Println("number is 2")

	}

}
