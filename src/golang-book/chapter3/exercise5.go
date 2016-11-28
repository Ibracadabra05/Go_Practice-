/*

	Program that converts from Fahrenheit into Celsius
	C = (F - 32) * 5/9

*/

package main

import "fmt"

func main() {

	var input float64

	fmt.Println("Please input a value in Fahrenheit: ")

	fmt.Scanf("%f", &input)

	output := (input - 32.0) * (5.0 / 9.0)

	fmt.Println("The value in celsius is:", output)
}
