/*
LESSON 10: Interfaces (expanded)

Overview
---------
An interface in Go is a set of method signatures. A concrete type satisfies
an interface implicitly by implementing those methods — there's no explicit
`implements` keyword.

Key concepts
------------
- Interface value: an interface value holds a (dynamic) type and a value.
	The zero value of an interface is `nil` (no type, no value).
- Empty interface: `interface{}` has zero methods and therefore is satisfied
	by every type. Use it when you need to accept any value (generic containers,
	logging helpers, etc.).
- Type assertion: `v, ok := i.(T)` extracts the concrete value of type T from
	interface `i` safely (ok=false when it doesn't hold T). The single-value
	form `x := i.(T)` will panic if `i` does not hold T.
- Type switch: `switch v := i.(type) { case int: ... }` lets you branch on the
	concrete type stored inside an interface.
- Prefer small interfaces: design interfaces that express minimal behavior
	(for example, `io.Reader` has a single `Read` method).

Practical tips for this lesson
------------------------------
- We'll implement a `Shape` interface (Area and Perimeter) and two concrete
	types (Circle, Rectangle). We'll also show how an interface can be used as
	a parameter, how to work with `interface{}`, how to assert types, and how
	type switches work.

This file includes short runnable examples demonstrating each idea. Read the
comments above each exercise for extra theory and notes.
*/

package main

import "fmt"

// Exercise 1: Basic interface definition
// TODO: Define a Shape interface with Area() and Perimeter() methods
type Shape interface {
	Area() float32
}

// Exercise 2: Implementing interfaces
// TODO: Create Circle and Rectangle structs that implement Shape interface
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}

// Exercise 3: Interface as function parameter
// TODO: Write a function that accepts Shape interface and prints area and perimeter
func PrintShape(s Shape) {
	var a float32 = s.Area()
	fmt.Println(a)
}

// Exercise 4: Multiple interface implementation
// TODO: Define a Drawable interface with Draw() method
type Drawable interface {
	Draw()
}

// TODO: Make Circle implement both Shape and Drawable
func (c Circle) Draw() {
	fmt.Println("Drawing a circle")
}

// Exercise 5: Empty interface
// TODO: Create a function that accepts interface{} and prints its type and value
func PrintAnything(value interface{}) {
	fmt.Println(value)
}

// Exercise 6: Type assertion
// TODO: Write functions that use type assertion to extract concrete types from interfaces
func PrintShapeType(s Shape) {
	_, isCircle := s.(Circle)
	if isCircle {
		fmt.Println("The shape is a circle")
	} else {
		fmt.Println("Is not a circle")
	}
}

// Exercise 7: Type switches
// TODO: Write a function that uses type switch to handle different types
type Animal interface {
	Walk()
	Sit()
}

type Human struct {
	Name string
}

type Dog struct {
	Name string
}

func (h Human) Walk() {
	fmt.Printf("The human %s is walking.\n", h.Name)
}
func (h Human) Sit() {
	fmt.Printf("The human %s is sitting.\n", h.Name)
}
func (d Dog) Walk() {
	fmt.Printf("The dog %s is walking.\n", d.Name)
}
func (d Dog) Sit() {
	fmt.Printf("The dog %s is sitting.\n", d.Name)
}
func Kiss(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("You kissed the dog %s.", v.Name)
	case Human:
		fmt.Printf("You kissed the human %s.", v.Name)
	}
}

// Exercise 8: Interface composition
// TODO: Create interfaces that embed other interfaces
type Bird interface {
	Fly()
}

type Swimmer interface {
	Swimm()
}

type SwimmerBird interface {
	Bird
	Swimmer
}

type Albatros struct {
}

func (a Albatros) Fly() {
	fmt.Println("The albatros can fly")
}
func (a Albatros) Swimm() {
	fmt.Println("The albatros can swimm")
}
func Describe(sb SwimmerBird) {
	sb.Fly()
	sb.Swimm()
}

func main() {
	var c Circle = Circle{
		Radius: 10,
	}
	fmt.Println(c.Area())
	c.Draw()
	fmt.Println("============================================================")
	PrintAnything(c)
	fmt.Println("============================================================")
	PrintShapeType(c)
	fmt.Println("============================================================")
	var d Dog = Dog{
		Name: "Marcus",
	}
	Kiss(d)
	d.Walk()
	fmt.Println("============================================================")
	var a Albatros = Albatros{}
	Describe(a)
	fmt.Println("============================================================")
	// Exercise 10: Interface slices
	// TODO: Create a slice of Animals interfaces with different concrete types
	var d1 Dog = Dog{
		Name: "Franco",
	}
	var d2 Dog = Dog{
		Name: "Pino",
	}
	var h1 Human = Human{
		Name: "Aleandro",
	}

	// TODO: Iterate and call methods on each
	animals := []Animal{d1, d2, h1}
	for _, animal := range animals {
		animal.Walk()
	}
	fmt.Println("============================================================")

	// Exercise 11: Nil interfaces and interface values
	// TODO: Understand nil interfaces vs interfaces with nil concrete values
	// TODO: Demonstrate the difference
	var i interface{} // nil interface, both type and value are nil
	var p *int = nil  // nil value but the type is still concrete (int)
	fmt.Printf("i == nil: %t\np == nil: %t\ni == p: %t\n", i == nil, p == nil, i == p)

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
