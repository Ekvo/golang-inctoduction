/*
Crete function 'work_with_new'
work_with_new()(*int) - define pointer and allocation memory for hin.
*/
package main

import "fmt"

func work_with_new() *int {
	i_ptr := new(int)
	*i_ptr = 5
	return i_ptr
}

func main() {
	fmt.Println(*work_with_new())
}
