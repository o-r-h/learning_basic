package exercise02

import "fmt"

func Examples_fmt() {
	//example for printing a string with a variable
	name := "John"
	age := 25
	price := 10.5
	isEstudent := true

	/*format string	 with:
	  %d for integer,
	  %s for string,
	  %f for float,
	  %t for boolean,
	*/
	fmt.Printf("Name: %s, Age: %d, Price: %f  is student %t\n", name, age, price, isEstudent)

	//format string with
	// %10s for string with 10 spaces
	// %-10s for string with 10 spaces
	// %10d for integer with 10 spaces
	// %10.3f for float with 10 spaces and 3 decimal places

	fmt.Printf("|%10s  example |/n", name)
	fmt.Printf("|%-10s |/n", name)
	fmt.Printf("|%10d |/n", age)
	fmt.Printf("|%10.3f n", price)

}
