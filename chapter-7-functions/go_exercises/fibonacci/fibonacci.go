/*
Create function 'fib' - fibonacci.
fib() - take one argument type uint

	and return n-value fibonacci

rules: fib(n) = fib(n-1) + fib(n-2);

	       fib(0) = 0;
		   fib(1) = 1;
*/
package main

import (
	"fmt"
	"strconv"
)

const loop int = 40 //my restrictions fibonachi is 40
const m_panic string = "!!!PANIC!!!"

func fib(val uint) (n_step uint) {
	switch val {
	case 0:
		n_step = 0
	case 1:
		n_step = 1
	default:
		n_step = fib(val-1) + fib(val-2)
	}
	return
}

func fibonacci_loop() {
	if loop > 40 {
		panic(m_panic + " bad loop, steps more than: 40 vs " + strconv.Itoa(loop))
	}
	for i := 0; i < loop; i++ {
		fmt.Println("Number:", i, " fibonacci equal:", fib(uint(i)))
	}
}

func catch_panic() { //catch block
	e_str := recover()
	if e_str != nil {
		fmt.Println(e_str)
	}
}

func main() {
	defer catch_panic()
	fibonacci_loop()
}
