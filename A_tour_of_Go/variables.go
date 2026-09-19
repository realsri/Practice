package main

import "fmt"

// declaring variables at package level
var c, java bool
var golang, python = true, "true" // intialized both the vars

// goal := "ஸ்ரீ" // error: non-declaration statement outside function body
var goal = "ஸ்ரீ" // same as 'var goal string = "ஸ்ரீ"'
// if an initializer is present, the type can be omitted; the variable will take the type of the initializer.

func main() {
	// declaring variables at function level
	var i int
	var j float32

	tamilRune := 'ஸ' //short assignment can be done without 'var' like this inside a function
	// same as - var tamilRune rune = 'ஸ'
	// rune is an alias for int32 and stores one Unicode code point.
	// can't set 'ஸ்ரீ' as a single rune as it is a sequence of multiple code points (ஸ +  ் + ர + ீ ) that combine visually into one displayed grapheme.
	// A string is a sequence of bytes, usually containing UTF-8-encoded text and may contain multiple runes.

	fmt.Println(i, j, c, java, python, golang, goal, tamilRune)

	// %c prints the rune as a character.
	fmt.Printf("Character: %c\n", tamilRune)

	// %U prints the Unicode code point.
	fmt.Printf("Unicode code point: %U\n", tamilRune)

	// %d prints the numeric code-point value.
	fmt.Printf("Numeric value: %d\n", tamilRune)

	// %T prints the Go type.
	fmt.Printf("Type: %T\n", tamilRune)
}
