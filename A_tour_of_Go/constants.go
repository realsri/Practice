package main

import "fmt"

const Pi = 3.14 // can be char/rune, string, bool or numeric values
// cannot declare using ':='
// An untyped constant takes the type needed by its context.

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 // stored as untyped int constant
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 // stored as untyped int constant
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "உலகம்"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))   // makes it work as int
	fmt.Println(needFloat(Small)) // makes it work as float64
	fmt.Println(needFloat(Big))   // makes it work as float64
	//fmt.Println(needInt(Big)) // can't fit into 'int'. An int can store at maximum a 64-bit integer, and sometimes less.
}
