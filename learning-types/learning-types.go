package main

import "fmt"

type duration int64

func main() {
	type Person struct {
		name  string
		age   int
		email string
	}

	jeff := Person{
		name:  "Jeff",
		age:   36,
		email: "jeffreylowy@gmail.com",
	}

	fmt.Println(jeff)
}
