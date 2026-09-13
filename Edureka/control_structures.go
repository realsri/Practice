package main

import "fmt"

func main() {
	defer lastrun() //runs at last
	firstrun()

	//panic("Something went wrong!") //stops normal execution

	fmt.Println(div(10, 2))
	fmt.Println(div(3, 0))
	fmt.Println(div(15, 5))
	//demPanic()
}

func firstrun() {
	fmt.Println("Running first")
}

func lastrun() {
	fmt.Println("Running last")
}

func div(num1, num2 int) int {
	//return num1 / num2 //panic: runtime error: integer divide by zero
	defer func() { //anonymous function //defer the function
		//fmt.Println(recover()) //call the recover function and print
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	res := num1 / num2
	fmt.Println("Function div executed successfully")
	return res
}

/*
func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
}*/
