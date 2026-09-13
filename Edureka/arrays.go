package main

import "fmt"

func main() {
	//var evennum [5]int

	evennum := [5]int{0, 2}
	for i := 2; i < 5; i++ {
		evennum[i] = i * 2
	}

	for i, value := range evennum {
		fmt.Println(i, value)
	}

	fmt.Println(evennum[2:])

	sliced := evennum[1:4]

	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)
	//allocated a size 10 array (can append untill 10 elements)
	// with the slice's length of 5
	// (i.e first 5 elements are 0 initialized - can access/modify directly like slice2[3])

	copy(slice2, sliced)

	fmt.Println(slice2)

	slice2 = append(slice2, 12, 34, 23, 34, 34, 34, 655)

	fmt.Println(slice2)

	slice3 := append(slice2, 45, 36, 27)

	fmt.Println(slice3)

}
