/*
LESSON 2: Variables and Constants

What you'll learn in this lesson:
- How to declare variables in different ways (explicit types vs type inference)
- How constants differ from variables and when to use them
- Basic built-in types and zero values

Core ideas explained for beginners:
- Declaration forms:
	* var x int = 10        // explicit type
	* y := 20               // short declaration, type inferred
	* var a, b = 1, "two"   // multiple declarations with inference

- Constants use the keyword const and must be compile-time values:
	const Pi = 3.14159

- Zero values: when you declare a variable without initializing it, Go gives it
	a sensible zero value (0 for numbers, "" for strings, false for bools, nil for pointers/slices/maps).

- Why types matter: Go is statically typed so operations require compatible types.
	Convert types explicitly with conversions: float64(i)

How to practice: try declaring the same value using different forms and print both
the value and its type using fmt.Printf("%v %T\n", value, value)
*/

package main

import "fmt"

func main() {
	// Exercise 1: Declare variables using different methods
	// TODO: Declare a string variable 'name' using var keyword
	var name string = "Hello!"
	// TODO: Declare an integer variable 'age' using short declaration (:=)
	age := 29
	// TODO: Declare a boolean variable 'isStudent' and set it to true
	isStudent := true
	// Use the variables in the format string so the program actually prints them.
	fmt.Printf("name: %s\nage: %d\nisStudent: %t\n", name, age, isStudent)

	// Exercise 2: Work with constants
	// TODO: Declare a constant 'pi' with value 3.14159
	const pi = 3.14159
	fmt.Printf("pi: %.2f\n", pi)
	// TODO: Declare a constant 'maxUsers' with value 1000
	const maxUsers = 1000
	fmt.Printf("maxUsers: %d\n", maxUsers)

	// Exercise 3: Multiple variable declarations
	// TODO: Declare variables 'firstName', 'lastName', and 'fullName' in a single var block
	var firstName, lastName, fullName string = "Aleandro", "Presta", "Aleandro Presta"
	fmt.Printf("firstName: %s\nlastName: %s\nfullName: %s\n", firstName, lastName, fullName)
	// TODO: es to firstName and lastName, then concatenate them into fullName
	firstName = "Aleandro"
	lastName = "Presta"
	fullName = firstName + " " + lastName
	fmt.Printf("fullName: %s\n", fullName)

	// Exercise 4: Type inference
	// TODO: Declare variables using := and let Go infer the types for:
	//       - A floating-point number
	//       - A boolean
	//       - A string
	a := 1.8
	b := false
	c := "Hello"
	// TODO: Print the values and their types using fmt.Printf with %T verb
	fmt.Printf("a: %T\nb: %T\nc: %T\n", a, b, c)

	// Exercise 5: Zero values
	// TODO: Declare variables without initial values and print them to see Go's zero values:
	//       - An integer
	//       - A string
	//       - A boolean
	var i int
	var j string
	var k bool
	fmt.Printf("i: %T\nj: %T\nk: %T\n", i, j, k)

	// Exercise 6: Variable reassignment
	// TODO: Declare a variable, print its value, change its value, and print again
	var d int = 5
	fmt.Printf("d: %d %T\n", d, d)
	d = 6
	fmt.Printf("d: %d %T\n", d, d)
}
