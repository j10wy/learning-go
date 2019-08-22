package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSeedValue() (x int64) {
	x = time.Now().Unix()
	return
}

func printRandomNumber(time int64) {
	rand.Seed(time)
	rNum := rand.Intn(100)
	fmt.Println("Number is", rNum)
}
