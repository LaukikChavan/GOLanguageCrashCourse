package main

import "fmt"

func main() {
	// Main Data Types in Go'
	/*
	* string
	* bool
	* int int8 int16 int32 int64
	* float32 float64
	* byte which is alise for int8
	* rune which is alise for int32
	* complex32 complex64
	* uint uint8 uint16 uint32 uint64 uintptr
	 */

	// variable by using var keyword
	// var name string = "Naruto"
	var age int16 = 20
	const isCool = true // can't assign false now

	//short hand method
	// name := "Naruto"
	// email := "example@some.com"
	pi := 3.14

	name, email := "Naruto", "example@some.com"

	fmt.Println(name, age, isCool, email, pi)
	fmt.Printf("name : %T \nage : %T \nisCool : %T \npi : %T", name, age, isCool, pi)
	fmt.Println()
}
