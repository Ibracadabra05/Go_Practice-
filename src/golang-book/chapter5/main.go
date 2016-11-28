/*
	Arrays, Slices, and Maps
*/

package main

import "fmt"

func main() {
	fmt.Println("Array...")
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	var total float64 = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
	fmt.Println("--------------")
	fmt.Println("Slice...")
	slicer()
	fmt.Println("Map...")
	mapper()
}

// slice practice

func slicer() {
	fmt.Println("Basic practice")
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[1:4]
	fmt.Println("x:", x)
	fmt.Println("--------------")

	fmt.Println("Underlying array practice")
	y := make([]int, 5, 10)
	fmt.Println("Length of slice y is: ", len(y))
	z := append(y, 1, 2, 3, 4, 5)
	fmt.Println("Length of underlying array is:", len(z))
	fmt.Println("underlying array looks like:", z)
	fmt.Println("y looks like:", y)
	fmt.Println("--------------")

	fmt.Println("Append practice")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice 1:", slice1, "slice 2:", slice2)
	fmt.Println("--------------")

	fmt.Println("Copy practice")
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3) // dst, src
	fmt.Println("slice 3:", slice3, "slice 4:", slice4)
	fmt.Println("--------------")
}

// map practice

func mapper() {
	fmt.Println("Basic way to initialize map")
	x := make(map[string]int)
	x["key"] = 10
	x["ellen"] = 7
	fmt.Println("x:", x)
	fmt.Println("x['key']:", x["key"])
	y := make(map[int]int)
	y[2] = 15
	fmt.Println("y:", y)
	fmt.Println("y[2]:", y[2])
	fmt.Println("--------------")

	fmt.Println("Deleting element from map")
	delete(x, "key")
	fmt.Println("x now looks like:", x)
	fmt.Println("--------------")

	fmt.Println("Checking if element exists in map")
	if name, ok := x["key"]; ok {
		fmt.Println("key still exists,", name, ok)
	} else {
		fmt.Println("key doesn't exist in x anymore as required,", name, ok)
	}
	fmt.Println("--------------")

	fmt.Println("Nested map practice")
	elements := make(map[string]map[string]string)
	elements["H"] = map[string]string{
		"name":  "Hydrogen",
		"state": "gas",
	}

	if el, ok := elements["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println("--------------")
}
