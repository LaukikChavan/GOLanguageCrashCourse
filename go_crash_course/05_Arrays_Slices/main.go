package main

import "fmt"

func main() {
	// Arrays
	var friuts [2]string

	// Assigning
	friuts[0] = "apples"
	friuts[1] = "oranges"

	// Declaration and Initialization
	games := [2]string{"Cricket", "Football"}

	fmt.Println(friuts)
	fmt.Println(games)

	// Sclice is like dynamic array it can increase its size and its declaration is preety similar to that array
	animes := []string{"Naruto", "Dragon ball", "Pokemon"}
	fmt.Println(animes)
}
