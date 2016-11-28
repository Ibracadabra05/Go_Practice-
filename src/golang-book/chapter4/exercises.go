/*
	1. Program that prints out all numbers between 1 and
	   100 that are divisible by 3

	2. Program that prints the numbers from 1 to 100, but for
	   multiples of three, print "Fizz" instead of the number,
	   and for the multiples of five, print "Buzz". For numbers
	   that are multiples of both three and five, print "FizzBuzz."
*/

package main

import "fmt"

func main() {
	fmt.Println("Running function one...")
	functionOne()
	fmt.Println("-----------------------")
	fmt.Println("Running function two...")
	functionTwo()
	fmt.Println("-----------------------")
}

func functionOne() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func functionTwo() {
	fizz := "Fizz"
	buzz := "Buzz"
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fizz + buzz)
		} else if i%3 == 0 {
			fmt.Println(fizz)
		} else if i%5 == 0 {
			fmt.Println(buzz)
		}
	}
}
