package main

import "fmt"

func main() {
	rect1 := rect{10, 20}
	rect2 := rect{w: 25, h: 15}

	fmt.Println(rect1.h, rect1.w)
	fmt.Println(rect2.h, rect2.w)

	fmt.Println(rect1.areamethod())
	fmt.Println(areafunc(rect2))
}

type rect struct {
	h float64
	w float64
}

func (r rect) areamethod() float64 {
	return r.h * r.w
}

func areafunc(r rect) float64 {
	return r.h * r.w
}
