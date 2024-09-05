/*
Print len type string "Hello, world"
then
Print simbol in string "Hello, world" with argiment 1 in []
then
Print type string "Hello, world" is sums of strings ("Hello" + "," + " " + "world")
--------------------------------------------------------------------------------------
Result is:
12
e
Hello, world
*/
package main

import "fmt"

func main() {
	fmt.Println(len("Hello, world"))
	fmt.Println("Hello, world"[1])
	fmt.Print("Hello" + "," + " " + "world")
}
