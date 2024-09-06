/*
Create function 'first_ptr'
first_ptr(*)(void) - take a pointer argument
*/
package main

import "fmt"

func first_ptr(i_ptr *int) { //take pointer and assign new value
	*i_ptr = 10
}

func main() {
	integer := 2
	fmt.Println("Bifore first_ptr(*) call: integer equal", integer)
	first_ptr(&integer)
	fmt.Println("After first_ptr(*) call: integer equal", integer)
}
