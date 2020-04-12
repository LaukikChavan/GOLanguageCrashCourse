package main

import "fmt"

func main() {
	// Declaration of maps
	// emails := make(map[string]string)

	// Assigning value to the maps
	// emails["Laukik"] = "laukikchavan101@gmail.com"
	// emails["example"] = "example@some.com"

	emails := map[string]string{"laukik": "laukik@gamil.com", "example": "example@gmail.com"}
	emails["apple"] = "apple@gmail.com"

	fmt.Println(len(emails))
	delete(emails, "laukik")

	fmt.Println(emails)

}
