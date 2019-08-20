package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getScore() {
	fmt.Print("Enter a score:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

func getFile() {
	file, err := os.Open("mytext.txt")
	// File will be nil
	if file != nil {
		// Error: open mytext: "no such file or directory"
		fmt.Print("err: ", err)
		log.Fatal(file)
	} else {
		fmt.Print("file: ", file.Name())
	}
}
