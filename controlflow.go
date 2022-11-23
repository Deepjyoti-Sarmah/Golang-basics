package main

import (
	"fmt"
)

func main() {
	// if i:=2 ; i == 3 {
	// 	fmt.Println("this is simple if statement")
	// } else if i == 2 {
	// 	fmt.Println("this is simple else if statement")
	// } else {
	// 	fmt.Println("this is simple else statement")
	// }

	// shoppingCart := make(map[string]int)
	// shoppingCart = map[string]int{
	// 	"Keyboard": 100,
	// 	"Mouse":    50,
	// 	"Laptop":   1000,
	// }

	// shoppingCart["Laptop"] = 1500
	// shoppingCart["Monitor"] = 1200

	// if _, ok := shoppingCart["Mic"] ; ok {
	// 	fmt.Println("Item exit in the shopping cart");
	// }

	// i :=20
	// j := 30

	// if i > 0  && j > 0{
	// 	fmt.Println("I is greater than zero")
	// }

	// switch i := 2 + 3; i {
	// case 1,3,5,7,9 :
	// 	fmt.Println("This is odd")
	// case 2,4,6,8 :
	// 	fmt.Println("This is even")
	// default:
	// 	fmt.Println("This is default")
	// }

	// i := 2 + 3
	// switch {
	// case i > 0:
	// 	fmt.Println("This is odd")
	// case i < 5:
	// 	fmt.Println("This is even")
	// default:
	// 	fmt.Println("This is default")
	// }

	// var i interface{} = 5.3
	// switch i.(type) {
	// case int:
	// 	fmt.Println("This is int type")
	// 	fmt.Println("This is int type")
	// 	break
	// 	fmt.Println("This is int type")
	// case float64:
	// 	fmt.Println("This is float type")
	// case string:
	// 	fmt.Println("This is string")
	// default:
	// 	fmt.Println("This is default")
	// }

	switch i := 2 + 3; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("This is odd")
		fallthrough
	case 2, 4, 6, 8:
		fmt.Println("This is even")
	default:
		fmt.Println("This is default")
	}

}
