/*
Create function 'swap'
swap(*,*)(void) - take two pointer and assign value with help dereferencing.
*/
package main

import "fmt"

var (
	x1 = 5
	x2 = 10
)

func my_swap(ptr_1 *int, ptr_2 *int) {
	ptr_tmp := ptr_1 //save first ptr
	*ptr_1 = *ptr_2
	*ptr_2 = *ptr_tmp
}

func main() {
	fmt.Println("Before swap: x1 equal", x1, ", x2 equal", x2)
	my_swap(&x1, &x2)
	fmt.Println("After swap: x1 equal", x1, ", x2 equal", x2)
}
