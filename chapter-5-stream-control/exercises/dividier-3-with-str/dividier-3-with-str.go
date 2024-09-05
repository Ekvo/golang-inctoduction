/*
In this exercises we print numbers that are divide on 3 and 5 and 15
with 3 var string
1. "Fizz" if %3
2. "Buzz" if %5
3. "FizzBuzz" if %15
*/
package main

import "fmt"

const loop int = 100 //count of numbers
const (
	fizz     int = 3
	buzz     int = 5
	fizzbuzz int = fizz * buzz
)

var (
	F string = "Fizz"
	B string = "Buzz"
)

func main() {
	for i := 0; i < loop; i++ {
		if i%fizzbuzz == 0 { // we can also use expression fizz*buzz, but it's expensive
			fmt.Println("Number: ", i, " is "+F+B)
		} else if i%fizz == 0 {
			fmt.Println("Number: ", i, " is "+F)
		} else if i%buzz == 0 {
			fmt.Println("Number: ", i, " is "+B)
		} else {
			fmt.Println("Number: ", i, " is common")
		}
	}
}
