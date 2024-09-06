/*Discription operator 'defer' (Delayed call)
-------------------------------------------------------
(it's necessary) Result is:
2-nd
1-st
3-d
*/
package main

import "fmt"

func first() {
	fmt.Println("1-st")
}
func second() {
	fmt.Println("2-nd")
}
func third() {
	fmt.Println("3-d")
}
func main() {
	second()
	defer third() //operator 'defer' move call third() at the end main()
	first()
}
