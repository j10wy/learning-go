package main

import (
	"fmt"
	"math/rand"
	"time"
)

func askQuestion() {
	fmt.Println("I've generated a random number. Can you guess it?")
}

func generateSeedValue() (x int64) {
	x = time.Now().Unix()
	return
}

func generateRandomNumber(time int64, maxVal int) (rNum int) {
	rand.Seed(time)
	rNum = rand.Intn(maxVal)
	return
}
