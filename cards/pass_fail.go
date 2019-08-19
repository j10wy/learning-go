package main

import (
	"bufio"
	"fmt"
	"os"
)

func getScore() {
	fmt.Print("Enter a score:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

func getFile() {
	file, err := os.Open("mytext")
	// File will be nil
	fmt.Print("file: ", file)
	// Error: open mytext: "no such file or directory"
	fmt.Print("err: ", err)
}
