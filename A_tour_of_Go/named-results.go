package main

import "fmt"

func split(sum int) (x, y int) { //return values may be named
	//treated as variables defined at the top of the function //notice that there are no ':' while assigning?
	x = sum * 4 / 9
	y = sum - x
	return //naked return //just do it only for short functions like this
	//you can also do 'return y, x' or 'return sum, x' as well (return 2 ints or mention nothing for default)
}

func main() {
	fmt.Println(split(17))
	a, b := split(17)
	fmt.Println("Ones and Tens:", a, b)
}
