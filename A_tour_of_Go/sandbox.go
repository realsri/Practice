package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Welcome to whatever it is")
	fmt.Print("\n")
	now := time.Now()
	fmt.Print("The current time is ", now)
	fmt.Println()
	fmt.Println(now.Month(), now.Day(), now.Year())
	fmt.Printf("%02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
}
