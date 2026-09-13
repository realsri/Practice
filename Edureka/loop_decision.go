package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 16
	for j < 10 {
		fmt.Println(j)
		j++
	}

	for ; ; j++ {
		fmt.Println(j)
		if j > 18 {
			break
		}
	}

	for {
		fmt.Println(j)
		j++
		if j > 20 {
			break
		} else {
			continue
		}
	}

	age := 17.5 //need not be integers //need not be constants //break is there already
	cage := 18.5

	switch age {
	case 16: //here age is float so converted to 16.0 and then compared
		fmt.Println("grow for two more years")
	case 17:
		fmt.Println("Just one more year")
	case 18:
		fmt.Println("You are right at it")
	//case j:
	//	fmt.Println("some number")
	case cage:
		fmt.Println("can compare with var of same type as well")
	default:
		fmt.Println("Get off")
	}

	fmt.Print("Thank you\n")
}
