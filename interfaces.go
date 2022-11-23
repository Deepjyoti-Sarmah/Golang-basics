//collection of behaviours
package main

import (
	"fmt"
	"math"
)

func main() {
	var r geometry = rectangle{width: 3, height: 4}
	fmt.Println(r.area(), r.perimeter())

	var c geometry = circle{5}
	fmt.Println(c.area(), c.perimeter())
}

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
