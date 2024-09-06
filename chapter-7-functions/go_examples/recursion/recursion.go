// Create function with name factorial (recursion method)
package main

import "fmt"

func factorial(val uint) uint {
	if val == 0 { //End recursion
		return 1
	}
	return val * factorial(val-1) //Next step with decrement
}
func main() {
	fmt.Println(factorial(10))
}
