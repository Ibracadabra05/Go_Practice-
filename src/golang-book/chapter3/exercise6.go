/*

	Program that converts from feet to meters

*/

package main

import "fmt"

const conv float64 = 0.3048

func main() {
	var input float64

	fmt.Println("Please enter value in feet for conversion to meters: ")

	fmt.Scanf("%f", &input)

	output := input * conv

	fmt.Println("Value in meters is:", output)
}
