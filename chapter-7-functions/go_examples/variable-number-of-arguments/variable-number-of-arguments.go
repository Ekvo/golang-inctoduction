// Create functiom with variable number of arguments
package main

import "fmt"

func variable_arguments(val ...int) (result int) { //find sum of arguments value
	for _, v := range val {
		result += v
	}
	return
}

func main() {
	array_ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(variable_arguments(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), //Print several values
		variable_arguments(array_ints...)) //Print array

}
