/*
'switch' with 'for' cycle
Print strings from "zero" to "nine"
and
end with help 'break'
*/
package main

import "fmt"

const loop int = 100         //number of cycle steps
const end_str string = "END" //for end
var str string               //for printing strings

func main() {
	for i := 0; i < loop; i++ {
		switch i {
		case 0:
			str = "zero"
		case 1:
			str = "one"
		case 2:
			str = "two"
		case 3:
			str = "three"
		case 4:
			str = "four"
		case 5:
			str = "five"
		case 6:
			str = "six"
		case 7:
			str = "seven"
		case 8:
			str = "eight"
		case 9:
			str = "nine"
		default:
			str = end_str
		}
		fmt.Println(str)
		if str == end_str {
			break
		}
	}
}
