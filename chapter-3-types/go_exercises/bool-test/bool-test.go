/*
test expression (true && false) || (false && true) || !(false && false)
-----------------------------------------------------------------------
Result is:
true
*/
package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}
