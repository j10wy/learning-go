package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	fmt.Println("Package", strings.Title("Main"), "verson", rand.Intn(100012132130))
	card := fmt.Sprintln("test", newCard2())

	fmt.Println(card)

}

func newCard2() string {
	return "Ace of spades"
}
