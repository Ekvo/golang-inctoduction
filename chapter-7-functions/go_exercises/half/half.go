/*
Create function 'half'
half - take an argument float64 and

	       devide on 2 and
	       return result and status even/odd source value
*/
package main

import "fmt"

func half(val int) (float64, bool) {
	even_odd_b := false
	if val%2 == 0 {
		even_odd_b = true
	}
	return float64(val) / 2.0, even_odd_b
}

func main() {
	fmt.Println(half(14))
}
