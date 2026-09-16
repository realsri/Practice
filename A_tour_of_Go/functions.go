package main

import "fmt"

// 2 params/args and 1 return value here
func add(x, y int) int { //same as func add(x int, y int) int {
	return x + y
}

// Refer https://go.dev/blog/declaration-syntax for more info
func main() {
	var a int = 40
	b := 60
	fmt.Printf("Let us add something (%d and %v) and learn functions to get %v\n", a, b, add(a, b))
}
