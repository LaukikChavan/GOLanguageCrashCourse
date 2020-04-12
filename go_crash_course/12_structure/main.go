package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver Method
func (this Person) greet() string {
	return "Hii! I am " + this.firstName + " " + this.lastName + " I am " + strconv.Itoa(this.age) + " Years old" + " from " + this.city + " Nice to Meet You."
}

// Refrence receiver Method
func (this *Person) hasBirthday() {
	this.age++
}

func main() {
	// Assignment of Data with labels
	//person1 := Person{firstName: "Naruto", lastName: "Uzumaki", city: "Konoha", gender: "M", age: 17}

	// Assignment of Data without labels
	person1 := Person{"Naruto", "Uzumaki", "Konoha", "M", 17}
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
