//Variables

package main

import "fmt"

const firstConstant string = "Ibracadabra"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	var x string
	var y string

	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "hello"
	y = "world"

	fmt.Println(x == y)

	//best way to assign variables
	z := 17
	fmt.Println(z)

	fmt.Println("a is", a)
	fmt.Println("b is", b)
	fmt.Println("c is", c)

	myfunc()
}

func myfunc() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
