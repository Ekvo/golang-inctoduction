/*
Distance conversion
feet to meters
*/
package main

import "fmt"

const cast_constant float64 = 0.3048

func main() {
	feet := 123.56
	meters := feet * cast_constant

	fmt.Println(meters)
}
