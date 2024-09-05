// In this exercises we print numbers that are divide on 3
package main

import "fmt"

const loop int = 100 //count of numbers
const divide int = 3

func main() {
	for i := 0; i < loop; i++ {
		if i%3 == 0 {
			fmt.Println(i, " is divide on ", divide)
		}
	}
}
