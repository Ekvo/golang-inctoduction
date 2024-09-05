/*
Temperature conversion
Fahrenheit to Celsius
*/
package main

import "fmt"

const water_freezing float64 = 32 //32 degrees Fahrenheit (°F)

func main() {
	Fahrenheit := 123.0
	Celsius := (Fahrenheit - water_freezing) * 5.0 / 9.0

	fmt.Println(Celsius)
}
