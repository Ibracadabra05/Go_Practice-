/*
	Functions
*/

package main

import "fmt"

func main() {
	xs := []float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(average(xs))
	fmt.Println(f())
	x, y := f1()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))

	//using slice on variadic function
	g := []int{1, 2, 3}
	fmt.Println(add(g...))

	//closure
	subtract := func(x, y int) int {
		return x - y
	}
	fmt.Println(subtract(5, 3))
	h := 8
	increment := func() int {
		h++
		return h
	}
	fmt.Println(increment())
	fmt.Println(increment())

	evenGenerator := func() func() uint {
		i := uint(0)
		return func() (ret uint) {
			ret = i
			i += 2
			return
		}
	}

	nextEven := evenGenerator()

	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
	fmt.Println(nextEven()) //6

	//recursion
	fmt.Println(factorial(2))

	//defer
	defer second()
	first()

	//panic and recover
	// defer func() {
	// 		str := recover()
	// 		fmt.Println(str)
	// 	}()
	// 	panic("PANIC")

	//pointer
	c := 48
	zero(&c)
	fmt.Println("c is", c)

	d := new(int)
	one(d)
	fmt.Println(*d)
}

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func average(xs []float64) float64 {
	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}

// named return values
func f() (r int) {
	r = 1
	return
}

// multiple return values
func f1() (int, int) {
	return 5, 6
}

// variadic function
func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}
