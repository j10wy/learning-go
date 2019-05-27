package main

import "fmt"

func main() {

	cards := []string{newCard(), "test"}

	for i, card := range cards {
		fmt.Println(i, card)
	}

}
