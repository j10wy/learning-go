package main

import (
	"fmt"
	"math/rand"
)

var (
	a = 10
	b = 14
)

func main() {
	fmt.Println("Henlo, fren!")
	fmt.Printf("Did you know there are %v hours in a day?\n", a+b)

	b++
	a = rand.Intn(1098700 / 1000)
	fmt.Println("Random number:", a)
}
