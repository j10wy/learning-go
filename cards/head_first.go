package main

import (
	"fmt"
	"time"
)

// PrintYear - print the exact time at run-time
func printYear() {
	var now = time.Now()
	var year = now.Year()
	fmt.Println(year)
}
