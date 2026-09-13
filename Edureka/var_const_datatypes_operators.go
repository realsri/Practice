package main

import "fmt"

func main() {
	const pi float64 = 3.14159265359 //constant or literal

	var a int = 5

	var b float32 = 4.32

	x, y := 8, 6

	var (
		varA = 2
		varB = 3
	) //declare multiple variables

	fmt.Println(varA, varB)

	fmt.Printf("%d %.1f\n", a, b)

	fmt.Printf("%T %T\n", a, b) //type

	fmt.Println(x, ",", y)

	//Arithmetic operators
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	var name string = "Sri"

	lame := "no"

	fmt.Println(name, len(name)) //string length

	fmt.Printf("%s \n", lame)

	fmt.Println(name + lame)

	var ans bool = true
	ques := false

	//Logical operators
	fmt.Printf("%t %t\n", ans, ques) //bool
	fmt.Println(ans && ques)
	fmt.Println(ans || ques)
	fmt.Println(!ans)

	//datatypes

	//Numeric - int (and uint) (8, 16, 32, 64) and float (float32 and float64)
	//String - string
	//Boolean (true and false)
	//Derived : Pointer, Array, Structure, Map, Interface

	//Relational operators
	// >, <, >=, <=, ==, !=

	fmt.Println(x)
	fmt.Println(&x)

	changeValue(&x)
	fmt.Printf("%d \n", x)

	fmt.Printf("%b \n", 25) //binary

	fmt.Printf("%c \n", 80) //ascii

	fmt.Printf("%x \n", 15) //hex

	fmt.Printf("%e \n", pi) //scientific notations

}

func changeValue(x *int) {
	*x = 7
}
