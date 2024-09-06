/*
Find lenght of  the slice
array is ["a","b","c","d","e","f"]
slice is [2:5]
------------------------------------------------------
result is:
Slice [ 2 : 5 ] of the array [a b c d e f] is [c d e]
*/
package main

import "fmt"

const array_len int = 6
const slice_begin int = 2
const slice_end int = 5

var array [array_len]string

func main() {
	array = [...]string{"a", "b", "c", "d", "e", "f"}
	slice := array[slice_begin:slice_end]
	fmt.Println("Slice [", slice_begin, ":", slice_end, "] of the array", array, "is", slice)

}
