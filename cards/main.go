package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	fmt.Println("Package", strings.Title("Main"), "verson", rand.Intn(100012132130))
	hand, cards := deal(newDeck(), 3)

	var a = []string{"One", "Two"}

	for i := range a {
		fmt.Println(a[i])
	}

	hand.print()
	cards.print()

}
