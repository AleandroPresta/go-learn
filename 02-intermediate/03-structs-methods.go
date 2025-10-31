/*
LESSON 9: Structs and Methods

Structs group related data together. They are the primary way to create
custom types in Go. Methods attach behavior to types.

Declaring structs:
type Person struct {
	Name string
	Age int
}

Methods:
func (p Person) Greet() string { ... }   // value receiver: works on a copy
func (p *Person) Birthday() { p.Age++ }  // pointer receiver: mutates the original

Key differences:
- Use pointer receivers when method needs to modify the receiver or to avoid copying large structs.
- Embedded structs implement composition: type Employee struct { Person; EmployeeID string }

Practice tip: design small structs (Point, Rectangle) and implement methods like Area and Translate.
*/

package main

import "fmt"

// Exercise 1: Basic struct definition
// TODO: Define a Person struct with Name (string), Age (int), and Email (string) fields

// Exercise 2: Struct with methods
// TODO: Define a Rectangle struct with Width and Height fields
// TODO: Add methods: Area(), Perimeter(), and IsSquare()

// Exercise 3: Struct with pointer receiver methods
// TODO: Define a Counter struct with a count field
// TODO: Add methods: Increment(), Decrement(), Reset(), Value()
// TODO: Use pointer receivers to modify the struct

// Exercise 4: Nested structs
// TODO: Define an Address struct with Street, City, Country fields
// TODO: Modify Person struct to include an Address field

// Exercise 5: Anonymous structs and embedded fields
// TODO: Create an anonymous struct for temporary data
// TODO: Define a struct with embedded fields (composition)

func main() {
	fmt.Println("============================================================")
	// Exercise 6: Struct initialization
	// TODO: Create Person instances using different initialization methods:
	//       - Zero value
	//       - Positional initialization
	//       - Named field initialization

	// Exercise 7: Struct operations
	// TODO: Access and modify struct fields
	// TODO: Copy structs and observe value semantics

	// Exercise 8: Method calls
	// TODO: Create Rectangle instances and call all methods
	// TODO: Create Counter instances and test increment/decrement

	// Exercise 9: Pointer vs value receivers
	// TODO: Demonstrate the difference between pointer and value receivers
	// TODO: Show when modifications persist and when they don't

	// Exercise 10: Struct comparison
	// TODO: Compare structs for equality
	// TODO: Understand when structs are comparable

	// Exercise 11: JSON-like operations (manual)
	// TODO: Create a function that converts a Person struct to a string representation
	// TODO: Create a function that creates a Person from string data

	// Exercise 12: Practical example
	// TODO: Design a simple bank account system with:
	//       - Account struct (AccountNumber, Balance, Owner)
	//       - Methods: Deposit(), Withdraw(), GetBalance(), Transfer()
	//       - Validation and error handling
}

// Implement all structs and methods here:
