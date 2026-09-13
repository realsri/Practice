package main

import (
	"fmt"
	"math"
)

func main() {
	rect1 := rect{20, 30}
	cir1 := circle{7}

	fmt.Println(getarea(rect1))
	fmt.Println(getarea(cir1))
}

type shape interface {
	area() float64
}

type rect struct {
	h, w float64
}

type circle struct {
	r float64
}

func (r1 rect) area() float64 {
	return r1.h * r1.w
}

func (c1 circle) area() float64 {
	return math.Pi * math.Pow(c1.r, 2)
}

func getarea(s shape) float64 {
	return s.area()
}
