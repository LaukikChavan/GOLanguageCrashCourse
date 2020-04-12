package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("%T \n%T \n", a, b)

	*b = 10

	fmt.Println(a)

}
