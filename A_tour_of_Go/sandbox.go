package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Welcome to whatever it is")
	fmt.Print("\n")
	// for new line

	now := time.Now()

	fmt.Print("The current time is ", now)
	// no space between arguments and no new line at the end

	fmt.Println()
	// for new line

	fmt.Println(now.Month(), now.Day(), now.Year())
	// print spaces between the arguments and a new line as well at the end

	fmt.Printf("%02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
	// formatted printing
	// %02d - format as min 2 decimal integers, use 0 to pad any empty spaces
}
