/*
Discription work functions 'panic' & 'recover';
-----------------------------------------------
panic - this function stop program in run time and generate committing an error;
recover - this function stop the 'panic" and return value passed to the 'panic' function;
*/
package main

import "fmt"

const m_panic string = "!!!PANIC!!!"

func catch_panic() {
	error_str := recover() //catch panic and copy message to error_str
	if error_str != nil {  //if panic is true
		fmt.Println(m_panic, error_str)
	}
}

func main() {
	defer catch_panic()
	panic("Exception") //SABOTAGE
}
