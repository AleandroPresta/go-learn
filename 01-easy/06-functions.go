/*
LESSON 6: Functions

Functions group logic into reusable units. In Go, functions are simple but
powerful: they can return multiple values and be passed around as values.

Syntax recap:
func name(param1 type1, param2 type2) (ret1 type1, ret2 type2) {
	// body
	return v1, v2
}

Key points:
- Multiple returns are idiomatic in Go (e.g., returning a value and an error).
- Named return values let you set results by name and use a bare return.
- Variadic functions accept a variable number of arguments: func sum(nums ...int) int
- Functions are values: assign them to variables or pass them as arguments.

Practice tip: Implement small functions first (add, multiply) then compose them
into higher-level operations. Try writing a function that returns another
function (a closure) to understand lexical scoping.
*/

package main

import "fmt"

// Exercise 1: Basic function
// TODO: Write a function that takes two integers and returns their sum

// Exercise 2: Function with multiple parameters
// TODO: Write a function that takes name (string) and age (int) and returns a greeting

// Exercise 3: Function with multiple return values
// TODO: Write a function that takes two numbers and returns both sum and product

// Exercise 4: Named return values
// TODO: Write a function that calculates area and perimeter of a rectangle using named returns

// Exercise 5: Variadic function
// TODO: Write a function that takes variable number of integers and returns their sum

func main() {
	// Exercise 6: Function calls
	// TODO: Call all the functions you created above and print their results
	
	// Exercise 7: Anonymous functions
	// TODO: Create an anonymous function that squares a number and assign it to a variable
	// TODO: Call the anonymous function
	
	// Exercise 8: Function as parameter
	// TODO: Write a function that takes another function as parameter
	// TODO: Create different operation functions (add, multiply) and pass them as arguments
	
	// Exercise 9: Closures
	// TODO: Create a function that returns a function (closure)
	// TODO: The returned function should have access to variables from the outer function
	
	// Exercise 10: Recursive function
	// TODO: Write a recursive function to calculate factorial
	// TODO: Write a recursive function to calculate Fibonacci numbers
}

// Helper functions area - implement your functions here:

// TODO: Implement all the functions for the exercises above