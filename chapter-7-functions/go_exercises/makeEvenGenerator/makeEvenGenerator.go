/*
Create function 'makeEvenGenerator'
makeEvenGenerator()  - like a struct with type func()

	and this generic 'odd' value and return him,
	take first odd number example (127,3,7)
	and then sum with 'even' value example (4,6,8),
	we get all time odd value.

-------------------------------------------------------
layout makeEvenGenerator(odd)  func(even) (return odd)
*/
package main

import (
	"fmt"
	"strconv"
)

const m_panic string = "!!!PANIC!!!"
const loop int = 100
const start_v int = 127 //v -value 'start_v' all time is odd or panic

func makeEvenGenerator(start int) func(tmp int) int {
	if start%2 == 0 {
		panic(m_panic + "start value no odd " + strconv.Itoa(start))
	}
	odd_int := start
	return func(tmp int) (result int) {
		if tmp%2 == 0 {
			result = odd_int
			odd_int += tmp
			return
		} else {
			panic(m_panic + "step value no even " + strconv.Itoa(tmp))
		}
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
	odds := makeEvenGenerator(start_v)
	for i := 0; i < loop; i++ {
		if i%2 == 0 {
			fmt.Println("Number", i, "with start value equal", start_v, "get number", odds(i))
		}
	}
}
