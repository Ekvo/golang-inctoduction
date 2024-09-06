// Create function with 'several' values to return
package main

import (
	"fmt"
	"strconv" //convert int to string for panic
)

const m_panic string = "!!!PANIC!!!"
const min_val int = -100000
const max_val int = 100000

var array [10]int

func several_values(arr []int) (min int, max int) { //find min and max values in range [min_val - max_val]
	if len(arr) == 0 {
		panic("Len array is equal 0")
	}
	min = max_val
	max = min_val
	for _, val := range arr {
		if val < min_val || val > max_val {
			panic("Data int array beyond " + strconv.Itoa(min_val) + " " + strconv.Itoa(max_val) + " is " + strconv.Itoa(val))
		}
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	return
}
func error_find() { //for call with defer
	e_str := recover()
	if e_str != nil {
		fmt.Println(m_panic, e_str)
	}
}
func main() {
	defer error_find()
	array = [...]int{100000, 245, -233, 3424, 545, 126, -907, -1928, 9019, 1}
	fmt.Println(several_values(array[:]))
}
