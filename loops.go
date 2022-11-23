package main

import "fmt"

func main() {

	// i, j := 0, 0
	// for i < 10 {
	// 	if i == 6 && j == 6 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Println(i, j)
	// 	i, j = i+1, j+1

	// }
	// fmt.Println(i, j)

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if i*j == 0 {
	// 			continue
	// 		}
	// 		fmt.Println(i * j)
	// 	}
	// }

	// a := []int{1, 2, 3, 4, 5, 6, 7}

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"Laptop":   1000,
	}

	shoppingCart["Laptop"] = 1500
	shoppingCart["Monitor"] = 1200

	for k:= range shoppingCart {
		fmt.Println(k)
	}
}
