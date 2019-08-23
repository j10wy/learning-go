package henlo

import (
	"fmt"
	"math/rand"
)

var (
	a = 10
	b = 14
)

// Doggo type def
type Doggo struct {
	Name  string
	Color string
	Age   int
}

// Wut is a function
func Wut() {
	fmt.Println("Henlo, fren!")
	fmt.Printf("Did you know there are %v hours in a day?\n", a+b)

	b++
	a = rand.Intn(1098700 / 1000)
	fmt.Println("Random number:", a)
}
