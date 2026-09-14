package main

import (
	"fmt"
	"math"
)

// good to import parenthesized as a group rather than using multiple import statements
// import "fmt"
// import "math"

func main() {
	fmt.Printf("Now I have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now I have %v problems again.\n", math.Sqrt(7))
	//can use %v in general (value in default format)
	//checkout https://pkg.go.dev/fmt for fmt documentation
}
