/*
Create function max
max - take an argument in format(val ...type)

	and return max value
*/
package main

import (
	"fmt"
	"strconv"
)

var array []int
var min_val int = -1000000

const m_panic string = "!!!PANIC!!!"

func max(args ...int) (result int) {
	result = min_val
	for _, val := range args {
		if val <= min_val { //if lower value than min_value (all date can be lower than min_val)
			//and we need mark (empty or no empty) container
			panic("bad value: " + strconv.Itoa(val))
		}
		if result < val {
			result = val
		}
	}
	if result == min_val { //if no data
		panic("func max([!!!]args ...int)")
	}
	return
}

func catch_panic() { //catch block
	e_str := recover()
	if e_str != nil {
		fmt.Println(m_panic, e_str)
	}
}

func main() {
	defer catch_panic()

	array = []int{1, 23, 4, 5, 456, -4569, 32459, 23942, 6, 7, -101010}
	fmt.Println("Max value in", array, "is", max(array...))
}
