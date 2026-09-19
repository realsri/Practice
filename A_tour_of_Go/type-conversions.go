package main

import (
	"fmt"
	"math"
)

func main() {
	k := 1
	x, y := 3, k      // here the type is inferred as int based on 3 ("untyped") and k ("typed")
	f1 := 3.142       // untyped so inferred as float64 depending on the precision of the constant
	g := 0.867 + 0.5i // untyped so inferred as complex128 depending on the precision of the constant
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//                        need explicit conversion
	var z uint = uint(f) // here as well as f is float64
	fmt.Println(k, x, y, f1, g, f, z)
	fmt.Printf("k type %T, x type %T, y type %T, f1 type %T, g type %T, f type %T, z type %T\n", k, x, y, f1, g, f, z)
}
