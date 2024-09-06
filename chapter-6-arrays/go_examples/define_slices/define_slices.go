// define slices - slices of the arrays
package main

import "fmt"

var slice1 = []int{1, 2, 3}
var slice2 = make([]int, 2)
var slice3 = append(slice1, 567)

//slice4:=append(slice1, 123,2345,124)//without 'var' dont work in global status

func main() {
	slice4 := append(slice3, 123, 2345, 124) //in local status work witout status 'var'
	slice5 := slice4[2 : len(slice4)-2]      //[3 567 123]
	copy(slice2, slice1)
	fmt.Println(
		"slice1 is", slice1,
		";slice2 is", slice2,
		";slice3 is", slice3,
		";slice4 is", slice4,
		";slice5 is", slice5)
}
