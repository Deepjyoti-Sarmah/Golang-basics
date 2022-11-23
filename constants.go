package main

import "fmt"

// const user string = "Admin" //global inside same package
// const User string = "Deep" //global available across packages
// const (
// 	User    string = "Admin"
// 	Product string = "Product"
// )

const (
	i = iota
	// j = iota
	// k = iota
	j
	k
)

const (
	a = iota + 1
	_
	_
	
	b
	c
	d
)

func main() {
	// const i int = 12
	// const j float32 = 3.4
	// const k string = "This is constant"
	// const l bool = true
	// var a int = 13
	// fmt.Println(i + a)

// 	fmt.Println(i)
// 	fmt.Println(j)
// 	fmt.Println(k)

fmt.Println(a,b,c,d);


}
