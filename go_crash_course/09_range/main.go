package main

import "fmt"

func main() {

	ids := []int{2, 3, 4, 5, 6, 2, 4, 6}

	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}

	emails := map[string]string{"laukik": "laukik@gamil.com", "example": "example@gmail.com"}

	for key, value := range emails {
		fmt.Printf("\n%s: %s\n", key, value)
	}

}
