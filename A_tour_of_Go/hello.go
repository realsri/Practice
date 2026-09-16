// Everything is a package in Go //Although groups of packages can be managed together as a module using go.mod
package main //main is the "executable package" //And everything else is a "Library package"
// All go files inside a single directory must share the exact same package name
// They can access each other's variables and functions freely without importing

// always start your programs with your package declaration

import "fmt"

// then import the required packages

// then continue with your list of variables, constants, functions in no strict order

// then main() - program's entry point

func main() {
	fmt.Println("Hello 👋, I am ஶ்ரீ")
}
