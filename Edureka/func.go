package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(factorial(5))

	x, y := 5, 10
	fmt.Println(x, "+", y, "=", add(x, y))

	fmt.Println(addmeup(10, 20, 30, 40, 50))

}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func addmeup(args ...int) int {
	sum := 0

	for _, val := range args {
		sum += val
	}
	return sum
}
