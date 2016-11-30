/* Parsing command line arguments */

package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 5, "the max value")
	flag.Parse()
	fmt.Println(rand.Intn(*maxp))
}
