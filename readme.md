# Go Notes

## Definitions

**_Concurrency_** - Easy-to-write software that does several operations simultaneously

## The WHY of Go

Go is a very consistent language. This is a good thing: you’ll often find you just know where to look in your project for a given piece of code, without having to think about it!

### Reasons go is awesome:

- Fast compilation
- Less cumbersome code
- Unused memory freed automatically (garbage collection)
- Easy-to-write software that does several operations simultaneously (concurrency)
- Good support for processors with multiple cores

## Packages

Every Go file starts with a package clause. A package is a collection of code that all does similar things, like formatting strings or drawing images. The package clause gives the name of the package that a file’s code will be a part of.

### The Main Package

Main is a special package which is required if the code is going to be run directly (usually from the terminal).

```go
package main

func main() {
	//...
}
```

When a Go program is run, it looks for a function named main and runs that first, which is why we named this function main. The main package always contains the main function.

## Imports

Go files almost always have one or more import statements. Each file needs to import other packages before its code can use the code those other packages contain. Loading all the Go code on your computer at once would result in a big, slow program, so instead you specify only the packages you need by importing them.

```go
// Single line import
import "fmt"

// import multiple packages with one import statement
import (
	"fmt"
	"strings"
)
```

## Types

Go is statically typed, which means that it knows what the types of your values are even before your program runs. If you use the wrong type of value in the wrong place, Go will let you know.
