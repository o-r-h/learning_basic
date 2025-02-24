package exercise03

import "fmt"

func Examples_cicles() {
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println("for loop", i)
	}

	//while loop
	i := 0
	for i < 5 {
		fmt.Println("while loop", i)
		i++
	}

	//do while loop
	i = 0
	for {
		fmt.Println("do while loop", i)
		i++
		if i >= 5 {
			break
		}
	}

	//for each loop
	numbers := []int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		fmt.Println("for each loop", i, number)
	}

	//for each loop
	//if you don't need the index, you can use the blank identifier _
	for _, number := range numbers {
		fmt.Println("for each loop", number)
	}

	//nested loops
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("nested loops", i, j)
		}
	}

	//break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println("break and continue", i)
	}

}
