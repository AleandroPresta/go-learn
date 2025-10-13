/*
LESSON 11: Error Handling

Go favors explicit error handling over exceptions. Functions that may fail
typically return an error as their final return value. The error type is an
interface with a single method: Error() string.

Common patterns:
- if err := doSomething(); err != nil { // handle error }
- Returning errors with context: fmt.Errorf("read failed: %w", err)
- Creating sentinel errors: var ErrNotFound = errors.New("not found")
- Custom error types implement Error() string to carry extra information.

Best practices:
- Check errors close to where they occur and return early if needed.
- Wrap errors when adding context using %w to allow unwrapping.
- Use sentinel errors carefully and prefer typed errors when you need structured data.

Practice tip: write functions that return (T, error) and build callers that handle
and propagate errors clearly.
*/

package main

import (
	"errors"
	"fmt"
)

// Exercise 1: Basic error handling
// TODO: Create a function that divides two numbers and returns an error for division by zero

// Exercise 2: Custom error types
// TODO: Define a custom error type for validation errors
// TODO: Implement the Error() method

// Exercise 3: Error wrapping and unwrapping
// TODO: Create a function that wraps errors with additional context
// TODO: Use errors.Unwrap() and errors.Is() to work with wrapped errors

// Exercise 4: Sentinel errors
// TODO: Define sentinel errors for common conditions (e.g., ErrNotFound, ErrInvalidInput)

// Exercise 5: Multiple error handling patterns
// TODO: Show different ways to handle errors in various scenarios

func main() {
	// Exercise 6: Error checking patterns
	// TODO: Demonstrate the standard error checking pattern
	// TODO: Show early returns for error handling
	
	// Exercise 7: Error aggregation
	// TODO: Create a function that can return multiple errors
	// TODO: Collect and report all errors instead of failing on the first
	
	// Exercise 8: Panic and recover
	// TODO: Demonstrate panic() and recover() (use sparingly)
	// TODO: Show when panic is appropriate vs returning errors
	
	// Exercise 9: Error context
	// TODO: Build error messages that provide useful debugging information
	// TODO: Include relevant values and operation context
	
	// Exercise 10: File operations with error handling
	// TODO: Read a file and handle various error conditions
	// TODO: Proper resource cleanup even when errors occur
	
	// Exercise 11: Validation with detailed errors
	// TODO: Create a struct validator that returns detailed error information
	// TODO: Handle multiple validation failures
	
	// Exercise 12: Error handling in HTTP-like scenarios
	// TODO: Simulate HTTP status codes with custom error types
	// TODO: Handle different error types differently
	
	// Exercise 13: Error logging and debugging
	// TODO: Create functions that log errors with different severity levels
	// TODO: Include stack trace information where helpful
	
	// Exercise 14: Retry mechanism with errors
	// TODO: Implement a retry function that attempts operations multiple times
	// TODO: Handle transient vs permanent errors differently
	
	// Exercise 15: Error handling best practices
	// TODO: Demonstrate guard clauses and early returns
	// TODO: Show proper error message formatting
	// TODO: Implement graceful degradation when possible
}

// Implement all error types and functions here: