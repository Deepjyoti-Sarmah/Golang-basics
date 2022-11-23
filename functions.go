package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	// msg := "hello"
	// writeMessage(&msg, "Hello")
	// fmt.Println(msg)

	// sum("sum:",1,2,3,4)

	// total := sum(1, 2, 3)
	// fmt.Println(total)

	// total,err := divide(3, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(total)

	// func () {
	// 	fmt.Println("this is anonymous function")
	// }()

	// fun := func () {
	// 	fmt.Println("this is anonymous function")
	// }

	// fun()

	rect := rectangle{
		width:  20,
		height: 30,
	}

	area := rect.area()
	fmt.Println(area)

}

// func writeMessage(msg *string, msg1 string) {
// 	*msg = "helo Deep"
// 	fmt.Println(*msg)
// }

// func sum(values ...int) int {

// 	total := 0
// 	for _, v := range values {
// 		total += v
// 	}
// 	// fmt.Println(total)
// 	return total
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 || b == 0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }
