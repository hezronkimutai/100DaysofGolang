// In Go (Golang), objects are created using structs,
//  which serve as the foundation for representing complex data structures.
// You can instantiate and work with structs to create objects with attributes and methods.

// Steps to Create and Use Objects in Go:
// Define a Struct:
// A struct is a composite data type that groups together variables under a single name.

// Instantiate the Struct:
// You can create an instance (object) of the struct using different techniques.

// Access and Modify Fields:
// You can access or modify the struct's fields using dot notation.

// Add Methods:
// You can associate methods with structs to implement behavior.

// Example: Creating and Using an Object in Go

package main

import "fmt"

// struct to represent a car
type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) DisplayInfo() {
	fmt.Printf("Car: %s %s (%d)\n", c.Brand, c.Model, c.Year)
}

func main() {
	car1 := Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
	}

	car2 := Car{"Honda", "Civic", 2019}

	car1.Year = 2010

	car1.DisplayInfo()
	car2.DisplayInfo()
}
