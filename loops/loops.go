package loops

import "fmt"

// Looper - a thing that loops
type Looper struct {
	MaxCount     int
	CurrentCount int
}

// Loop - A looping method
func (l Looper) Loop() {
	for x := 0; x <= l.MaxCount; x++ {
		fmt.Println("X:", x)
	}
}
