/*
Find the small value in array

	 []int{
	    48,96,86,68,
	    57,82,63,70,
	    37,34,83,27,
	    19,97,9,17,
	}

---------------------------------------------------------------
Result is:
The smallest value in array [48 96 86 68 57 82 63 70 37 34 83 27 19 97 9 17] is 9
*/
package main

import "fmt"

const square_side_matrix int = 4

var array [square_side_matrix * square_side_matrix]int //lenght 16

func main() {
	array = [...]int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := 100
	for _, val := range array {
		if smallest > val {
			smallest = val
		}
	}
	fmt.Println("The smallest value in array", array, "is", smallest)
}
