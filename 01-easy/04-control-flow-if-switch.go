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
	
	// Exercise 2: If with initialization
	// TODO: Use the pattern "if x := someValue; condition" to check if a number is even
	
	// Exercise 3: Nested if statements
	// TODO: Check age ranges: child (< 13), teenager (13-19), adult (20-64), senior (65+)
	
	// Exercise 4: Basic switch statement
	// TODO: Create a switch statement that prints the day of the week based on a number (1-7)
	
	// Exercise 5: Switch with multiple cases
	// TODO: Create a switch for grades: A,B (Excellent), C,D (Good), F (Fail)
	
	// Exercise 6: Switch with expressions
	// TODO: Use switch to categorize numbers: negative, zero, small positive (1-10), large positive (>10)
	
	// Exercise 7: Switch without expression (replaces if-else chain)
	// TODO: Use switch true with conditions to check temperature ranges:
	//       - Freezing (< 0°C)
	//       - Cold (0-15°C)
	//       - Mild (16-25°C)
	//       - Hot (> 25°C)
	
	// Exercise 8: Logical operators in conditions
	// TODO: Check if a year is a leap year using if statement
	//       (divisible by 4 AND (not divisible by 100 OR divisible by 400))
}