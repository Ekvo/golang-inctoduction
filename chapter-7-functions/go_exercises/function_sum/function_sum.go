/*
Create function with name 'sum'
sum - take an argumetn array or slice and

	find sum of container elements
*/
package main

import "fmt"

var array_int [10]int

func sum(arr []int) (result int) {
	for _, val := range arr {
		result += val
	}
	return
}
func main() {
	array_int = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Summa of array elements", array_int, "is", sum(array_int[:]))
}
