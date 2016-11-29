/*
Structs and Interfaces as well as exercise for chapter 7
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//exercise
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

//exercise
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	c := &Circle{x: 0, y: 0, r: 5}
	r := &Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}

	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}
