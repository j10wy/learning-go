package main

import (
	"fmt"

	"github.com/jeffreylowy/intro-to-go/henlo"
)

func main() {
	//randNum := generateRandomNumber(generateSeedValue(), 100)
	askQuestion()
	henlo.Wut()

	kals := henlo.Doggo{
		Name:  "Kali",
		Color: "gray and white",
		Age:   8,
	}

	k := fmt.Sprintf("My dog, %s, is %s. She is %v years old", kals.Name, kals.Color, kals.Age)
	fmt.Println(k)
}
