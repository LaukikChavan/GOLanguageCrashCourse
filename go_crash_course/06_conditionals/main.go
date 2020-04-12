package main

import "fmt"

func main() {
	x := 10
	y := 15

	if x >= y {
		fmt.Printf("%d is Greater than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is neither Red nor Blue")
	}

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is neither Red nor Blue")
	}

}
