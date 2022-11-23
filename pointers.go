package main

import "fmt"

type Foo struct {
	bar int
}

func main() {
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a)
	// fmt.Println(*b)

	var foo *Foo
	fmt.Println(foo)

	foo = new(Foo)
	// (*foo).bar = 45
	// fmt.Println((*foo).bar)
	foo.bar = 45
	fmt.Println(foo.bar)
}
