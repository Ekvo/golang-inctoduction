// Define function inside function is 'Closures'
package main

import "fmt"

var (
	a = 123
	b = 456
	c = 789
)

func make_func_inside(start int) func(tmp int) int { //value 'start' not change
	val := 0
	return func(tmp int) (result int) { //value 'tmp' need every call add manually
		result = val + tmp
		val += start
		return
	}
}
func main() {
	val := func() (result int) { //sum global values
		result = a + b + c
		return
	}
	val2 := func(values ...int) (result int) { //sum several values
		for _, v := range values {
			result += v
		}
		return
	}
	val3 := make_func_inside(123)
	fmt.Println(val(), //Print 'Closures' wit name val
		val2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), //Print 'Closures' wit name val2
		val3(1), val3(2), val3(3))           //Print 'Closures' wit name val3 - 3 times
}
