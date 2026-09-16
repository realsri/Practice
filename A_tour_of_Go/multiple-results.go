package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("abraka", "dabraka")
	fmt.Println("Swapped:", a, b)
	fmt.Println(swap("abraka", "dabraka"))
	//fmt.Println("Swapped:", swap("abraka", "dabraka")) //Error
	//Go does not allow a multi-valued function call to be used as an argument alongside other arguments.
}
