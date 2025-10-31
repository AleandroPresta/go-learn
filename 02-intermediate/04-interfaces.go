/*
LESSON 10: Interfaces

Interfaces specify behavior by listing method signatures. Any type that
implements those methods satisfies the interface implicitly—there's no
implements keyword.

Example:
type Stringer interface { String() string }

Important ideas:
- The empty interface interface{} can hold any value—useful for generic containers.
- Type assertions extract the concrete value: v, ok := i.(T)
- Type switches let you branch on the concrete dynamic type inside an interface.
- Prefer small interfaces (io.Reader is a classic single-method interface).

Practice tip: implement a Shape interface with Area() and Perimeter() and write
functions that accept Shape to operate on different concrete types.
*/

package main

import "fmt"

// Exercise 1: Basic interface definition
// TODO: Define a Shape interface with Area() and Perimeter() methods

// Exercise 2: Implementing interfaces
// TODO: Create Circle and Rectangle structs that implement Shape interface

// Exercise 3: Interface as function parameter
// TODO: Write a function that accepts Shape interface and prints area and perimeter

// Exercise 4: Multiple interface implementation
// TODO: Define a Drawable interface with Draw() method
// TODO: Make Circle implement both Shape and Drawable

// Exercise 5: Empty interface
// TODO: Create a function that accepts interface{} and prints its type and value

// Exercise 6: Type assertion
// TODO: Write functions that use type assertion to extract concrete types from interfaces

// Exercise 7: Type switches
// TODO: Write a function that uses type switch to handle different types

// Exercise 8: Interface composition
// TODO: Create interfaces that embed other interfaces

func main() {
	fmt.Println("============================================================")
	// Exercise 9: Interface usage
	// TODO: Create instances of different shapes and use them through interface

	// Exercise 10: Interface slices
	// TODO: Create a slice of Shape interfaces with different concrete types
	// TODO: Iterate and call methods on each

	// Exercise 11: Nil interfaces and interface values
	// TODO: Understand nil interfaces vs interfaces with nil concrete values
	// TODO: Demonstrate the difference

	// Exercise 12: Interface assertion patterns
	// TODO: Implement safe type assertion with the comma ok idiom
	// TODO: Handle cases where assertion fails

	// Exercise 13: Practical example - Plugin system
	// TODO: Design a simple plugin system using interfaces:
	//       - Plugin interface with Name() and Execute() methods
	//       - Multiple plugin implementations
	//       - Plugin manager that runs all plugins

	// Exercise 14: Standard library interfaces
	// TODO: Implement fmt.Stringer interface for your custom types
	// TODO: Understand how this integrates with fmt package

	// Exercise 15: Interface best practices
	// TODO: Create small, focused interfaces
	// TODO: Demonstrate accepting interfaces and returning concrete types
}

// Implement all interfaces and types here:
