package main

import "fmt"

/*
Write a function that takes an integer and halves it and returns true if
it was even or false if it was odd. For example, half(1) should return
(0, false) and half(2) should return (1, true)
*/

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

/*
Write a function with one variadic parameter that finds the greatest number
in a list of numbers
*/

func findGreatest(args ...int) (ret int) {
	temp := 0
	for _, val := range args {
		if val > temp {
			temp = val
		}
	}
	ret = temp
	return
}

/*
Write a makeOddGenerator function that generates odd numbers.
*/

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

/*
Write a recursive function that can find fib(n).
*/

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*
Write a program that can swap two integers.
*/

func swap(xPtr, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr
}

func main() {
	x, y := half(1)
	w, z := half(2)
	fmt.Println(x, y, w, z)
	fmt.Println(findGreatest(6, 8, 100, 7, 9))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(fib(5))
	fmt.Println(fib(4))

	a, b := 5, 6
	swap(&a, &b)
	fmt.Println("a =", a, "b =", b)
}
