package main

import "fmt"

func main() {

	// var amount [3]int = [3]int{10, 20, 30}
	// amt := [3]int{30, 40, 50}
	// amt := [...]int{30, 40, 50, 60, 70, 80}

	// a := amount
	// a[0] = 51

	// fmt.Printf("Amounts : %v\n", amount)
	// fmt.Printf("A:%v\n", a)

	// b := &amount
	// b[0] = 51

	// fmt.Printf("Amounts : %v\n", amount)
	// fmt.Printf("B:%v\n", b)

	// fmt.Printf("Amt : %v\n ", amt)
	// fmt.Printf("%v \n ", len(amt))

	// a := [...]int {1,2,3,4,5,6,7,8,9,10};
	// b := a[1: len(a)-1]
	// c := a[2:]
	// d := a[:5]
	// e := a[2:7]

	// fmt.Printf("A: %v\n", a)
	// fmt.Printf("B: %v\n", b)
	// fmt.Printf("C: %v\n", c)
	// fmt.Printf("D: %v\n", d)
	// fmt.Printf("E: %v\n", e)

	// var identityMatrix [3][3]int = [3][3]int {
	// 	[3]int {1,0,0},
	// 	[3]int {0,1,0},
	// 	[3]int {0,0,1},
	// }
	// identityMatrix[1][2] = 7;
	// fmt.Println(identityMatrix)

	//slices
	// var a[]int = []int{1,2,3}
	// var b []int = a

	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// fmt.Println(b)

	// b[0] = 5

	// fmt.Println(a)
	// fmt.Println(b)

	// var a[]int = make([]int,3, 10)

	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	var a []int = []int{1, 2, 3}
	var b []int = append(a, 5)
	var c []int = append(b, a...)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
