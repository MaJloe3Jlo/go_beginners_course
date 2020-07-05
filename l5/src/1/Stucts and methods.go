package main

import (
	"fmt"	
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area1() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) values() {
	c.x, c.y, c.r = 1, 2, 3
}

func (c *Circle) valuesByPtr() {
	c.x, c.y, c.r = 1, 2, 3
}

func main() {
	var c Circle
	c.values()
	fmt.Print(c, " - Circle after values c \n")
	c.valuesByPtr()
	fmt.Print(c, " - circle after values by pointer\n")
	c2 := Circle{1,2,3}
	c3 := new(Circle)
	fmt.Printf("%T, %T, %T \n", c, c2, c3)
	c.x = 4
	fmt.Printf("%v, %v \n", c2.area(), c2.area1())	
	fmt.Printf("%f", c.x)
}
