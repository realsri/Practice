package main

//import paths
import (
	"fmt"
	r "math/rand" //this package comprises files that begins with 'package rand'
)

func main() {
	fmt.Println("A random number:", r.Intn(15))
	/*
	 * Can use aliases as above;
	 * 'import "math"' and then 'math.rand' access wont work because
	 * 'math/rand' is a separate package ('rand' inside 'math' folder),
	 * not a 'rand' member inside the 'math' package.
	 */
}
