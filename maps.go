package main

import "fmt"

func main() {

	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"Laptop":   1000,
	}

	shoppingCart["Laptop"] = 1500
	shoppingCart["Monitor"] = 1200

	fmt.Println(shoppingCart)
	delete(shoppingCart, "Monitor")
	fmt.Println(shoppingCart)

	// sc := shoppingCart
	
	// fmt.Println(shoppingCart)
	// fmt.Println(sc)

	// sc["Monitor"] = 1350
	// fmt.Println(shoppingCart)
	// fmt.Println(sc)

	// cart, ok  := shoppingCart["Mobile"]
	// fmt.Println(cart, ok)

	// _, ok  := shoppingCart["Mobile"]
	// fmt.Println( ok)

	

	

}
