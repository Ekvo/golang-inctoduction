// cycle for
package main

import "fmt"

const loop int = 100

/*
func main() {      //ancestor layout
	i := 0
	for i < loop {
		fmt.Println(i)
		i++
	}
}
*/
func main() {
	for i := 0; i < loop; i++ {
		if i == 0 {
			fmt.Println(i, "no divisible")
		} else if i > 2 && i%2 == 0 {
			fmt.Println(i, " is divisible by 2")
		} else if i > 3 && i%3 == 0 {
			fmt.Println(i, " is divisible by 3")
		} else if i > 5 && i%5 == 0 {
			fmt.Println(i, " is divisible by 5")
		} else if i > 7 && i%7 == 0 {
			fmt.Println(i, " is divisible by 7")
		} else {
			fmt.Println(i, " is divisible by himself")
		}
	}
}
