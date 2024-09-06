/*
Create function with name 'avarage'
func 'avarage' take array or slice and then return avarage value
*/
package main

import "fmt"

const m_panic string = "!!!PANIC!!!" //message

func avarage(arr []float64) (total float64) { //define default 'total' with type float64 he store 0.0
	if len(arr) == 0 { //exception
		panic("Len arr equal zero: divide by zero!")
	}
	for _, val := range arr {
		total += val
	}
	total /= float64(len(arr)) //find avarage
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
	//var arr []float64 //for test error_find()
	arr := []float64{1.45, 67.8, 456.1, -36.5, 657.7, 903.9}
	fmt.Println("Avarage of", arr, "is", avarage(arr))
}
