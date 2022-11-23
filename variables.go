package main

import (
	"fmt"
	"strconv"
)

var i = 32

func main() {

	// var i int
	// i = 12
	// var i int = 12
	
	var i = 12
	// j := 12
	var j float32 = float32(i)
	j = 3.4

	k := 3.4

	// var foo string = string(i)
	var foo string = strconv.Itoa(i)

	var bar bool = true

	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)
	fmt.Printf("%v , %T\n", foo, foo)
	fmt.Printf("%v , %T\n", bar, bar)

}