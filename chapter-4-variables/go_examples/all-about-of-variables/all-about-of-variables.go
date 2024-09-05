package main

import "fmt"

var g_str string = "Global world" //global var
const g_x int = 3                 //global var
var (                             //global vars
	a = 1
	b = 2
	c = 3
)

func some_f() {
	fmt.Println(g_str) //Print - Global world
}

func main() {
	fmt.Println(g_str) //Print - Global world
	some_f()

	var str_first string = "Hello, world" //local var in main
	var str_second string = str_first     //local var in main
	var str_third string                  //local var in main
	str_second = "and have a nice day"
	str_third = str_first + " " + str_second + " and sunshine in the sky."
	fmt.Println(str_third)               //Print - Hello, world and have a nice day and sunshine in the sky.
	fmt.Println(str_first == str_second) //Print - false

	without_var := "Hello my friend! Glad to see you."
	fmt.Println(without_var) //Print - Hello my friend! Glad to see you.

	x := 123       // local var in main
	fmt.Println(x) //Print - 123

	fmt.Println(g_x)     //Print - 3
	fmt.Println(a, b, c) //Print - 1 2 3
}
