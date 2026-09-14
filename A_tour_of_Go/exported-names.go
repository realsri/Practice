package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) //It is 'Pi', not 'pi'
	//'Pi' starting with Capital letter is exported name from the 'math' package
	//unexported names are not accessible from outside the package
}
