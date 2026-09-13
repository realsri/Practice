package main

import "fmt"

func main() {
	studentage := make(map[string]int)
	////////////////////////key   //value
	studentage["Sri"] = 25
	studentage["dha"] = 52
	studentage["ran"] = 0

	fmt.Println(studentage)

	fmt.Println(studentage["dha"])
	fmt.Println(studentage["Dha"])

	fmt.Println(len(studentage))

	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"name": "Clark Kent",
			"city": "Metropolis",
		},
		"Batman": map[string]string{
			"name": "Bruce Wayne",
			"city": "Gotham",
		},
	}
	fmt.Println(superhero)

	fmt.Println(superhero["Superman"], superhero["Batman"])

	fmt.Println(superhero["Batman"]["name"])

	//safer to check if the key exists and then access like below
	if temp, status := superhero["Superman"]; status {
		fmt.Println(temp["name"], temp["city"])
	}
}
