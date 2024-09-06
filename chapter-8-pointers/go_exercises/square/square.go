/*
Find solution for function 'square'
square(*)(void) - find square of value
*/
package main

import "fmt"

func square(f_ptr *float64) {
	*f_ptr = *f_ptr * *f_ptr
}

func main() {
	value := 10.5
	fmt.Println("value", value)
	square(&value)
	fmt.Println("square value", value)
}
