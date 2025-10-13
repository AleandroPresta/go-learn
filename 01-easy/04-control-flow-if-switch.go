/*
LESSON 4: Control Flow - If Statements and Switch

Control flow decides which code runs under which conditions. Go uses standard
if and switch constructs. Read the quick guide below, then try the exercises.

If statements:
- Basic form: if condition { ... }
- With else: if condition { ... } else { ... }
- With initialization: if v := compute(); v > 0 { ... }

Important points:
- Conditions must evaluate to a boolean (no truthy/falsy values like some languages).
- Short variable declaration in the if header limits variable scope to the if/else blocks.

Switch statements:
- switch value { case x: ... case y: ... default: ... }
- Cases don't fall through by default; use the fallthrough keyword to continue to the next case.
- You can write switch with no expression (switch { case cond1: ... }) as a cleaner if-else chain.

Practical advice:
- Prefer small, clear conditions. Use helper functions for complex predicates.
- Use switch for multiple branches over the same expression.
*/

package main

import "fmt"

func main() {
	// Exercise 1: Basic if statements
	// TODO: Write an if statement that checks if a number is positive, negative, or zero
	var i int = 10
	if i > 0 {
		fmt.Printf("%d is positive!\n", i)
	} else if i < 0 {
		fmt.Printf("%d is negative!\n", i)
	} else {
		fmt.Printf("%d is zero!\n", i)
	} // Exercise 2: If with initialization
	// TODO: Use the pattern "if x := someValue; condition" to check if a number is even
	if x := 11; x%2 == 0 {
		fmt.Println("x is even!")
	} else {
		fmt.Println("x is odd!")
	}

	// Exercise 3: Nested if statements
	// TODO: Check age ranges: child (< 13), teenager (13-19), adult (20-64), senior (65+)
	if x := 29; x <= 13 {
		fmt.Println("x is a child")
	} else if x > 13 && x <= 19 {
		fmt.Println("x is a teenager")
	} else if x > 19 && x <= 64 {
		fmt.Println("x is an adult")
	} else if x > 64 {
		fmt.Println("x is a senior")
	}
	// Exercise 4: Basic switch statement
	// TODO: Create a switch statement that prints the day of the week based on a number (1-7)
	var day int = 4
	switch day {
	case 1:
		fmt.Println("Today is monday!")
	case 2:
		fmt.Println("Today is tuesday!")
	case 3:
		fmt.Println("Today is wednesday!")
	case 4:
		fmt.Println("Today is thursday!")
	case 5:
		fmt.Println("Today is friday!")
	case 6:
		fmt.Println("Today is saturday!")
	case 7:
		fmt.Println("Today is sunday!")
	}
}
