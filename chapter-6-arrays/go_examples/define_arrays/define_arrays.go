// define arrays
package main

import "fmt"

var arr [7]int // array len 7 contain [0 0 0 0 0 0 0]
func main() {
	//arr=[...]int{98, 93, 77, 82, 83,56,78}//first way to add date to array
	/*
		arr = [...]int{//second way to add date to array
			98,
			93,
			77,
			82,
			83,
			56,
			78, // attention to comma in 78, it's necessary (if we use //78,)
		}
	*/
	len_arr := len(arr)            //int
	for i := 0; i < len_arr; i++ { //third way to add date to array
		arr[i] = i*i + 10
	}
	total := 0.0              //float64
	for _, val := range arr { //find the summ of all elements arr[]
		total += float64(val)
	}
	fmt.Println("Avarege number in array with len equal:", len_arr, "contain", arr, "is", total/float64(len_arr))
}
