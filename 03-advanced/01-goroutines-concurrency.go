/*
LESSON 13: Goroutines and Concurrency

Goroutines are lightweight concurrent functions managed by the Go runtime.
They enable writing concurrent programs without the heavy cost of OS threads.

Key points:
- Start a goroutine with: go someFunction()
- The main goroutine runs main(); when it returns, the program exits even if other goroutines are running.
- Use channels or synchronization primitives (sync.WaitGroup, mutexes) to coordinate goroutines.
- Always design a way to stop goroutines (context.Context, done channels) to prevent leaks.

Debugging tips:
- Race conditions can be detected with: go run -race .
- Prefer message passing (channels) to shared mutable state when possible.

Practice tip: implement small worker pools and use WaitGroup to ensure all workers finish.
*/

package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Basic goroutine creation
	// TODO: Create a simple goroutine that prints numbers
	// TODO: Use time.Sleep to see concurrent execution
	fmt.Println("============================================================")
	fmt.Println("=== EXERCISE 1: Basic Goroutines ===")

	// Exercise 2: Multiple goroutines
	// TODO: Create multiple goroutines that run concurrently
	// TODO: Observe the execution order (non-deterministic)
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 2: Multiple Goroutines ===")

	// Exercise 3: Anonymous goroutines
	// TODO: Create goroutines using anonymous functions
	// TODO: Pass parameters to anonymous goroutines
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 3: Anonymous Goroutines ===")

	// Exercise 4: WaitGroup synchronization
	// TODO: Use sync.WaitGroup to wait for goroutines to complete
	// TODO: Ensure main doesn't exit before goroutines finish
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 4: WaitGroup ===")
	waitGroupExample()

	// Exercise 5: Race conditions
	// TODO: Create a race condition with shared variables
	// TODO: Demonstrate the problem of concurrent access
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 5: Race Conditions ===")
	raceConditionExample()

	// Exercise 6: Mutex for synchronization
	// TODO: Fix the race condition using sync.Mutex
	// TODO: Protect shared resources with locks
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 6: Mutex ===")
	mutexExample()

	// Exercise 7: Read-Write mutex
	// TODO: Use sync.RWMutex for scenarios with many readers, few writers
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 7: RWMutex ===")
	rwMutexExample()

	// Exercise 8: Atomic operations
	// TODO: Use sync/atomic package for simple atomic operations
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 8: Atomic Operations ===")
	atomicExample()

	// Exercise 9: Worker pool pattern
	// TODO: Implement a worker pool with goroutines
	// TODO: Distribute work among multiple workers
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 9: Worker Pool ===")
	workerPoolExample()

	// Exercise 10: Context for cancellation
	// TODO: Use context.Context to cancel goroutines
	// TODO: Implement graceful shutdown
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 10: Context Cancellation ===")
	contextExample()
}

func waitGroupExample() {
	// TODO: Implement WaitGroup example
	fmt.Println("TODO: Implement WaitGroup synchronization")
}

func raceConditionExample() {
	// TODO: Create and demonstrate a race condition
	fmt.Println("TODO: Demonstrate race condition")
}

func mutexExample() {
	// TODO: Fix race condition using Mutex
	fmt.Println("TODO: Use Mutex to fix race condition")
}

func rwMutexExample() {
	// TODO: Implement RWMutex example
	fmt.Println("TODO: Implement RWMutex example")
}

func atomicExample() {
	// TODO: Use atomic operations
	fmt.Println("TODO: Use atomic operations")
}

func workerPoolExample() {
	// TODO: Implement worker pool pattern
	fmt.Println("TODO: Implement worker pool")
}

func contextExample() {
	// TODO: Use context for cancellation
	fmt.Println("TODO: Use context for cancellation")
}

// Additional helper functions for exercises:

// TODO: Implement a counter that multiple goroutines increment
// TODO: Implement a concurrent-safe cache
// TODO: Implement a rate limiter using goroutines
// TODO: Implement a pipeline pattern with goroutines
// TODO: Implement a fan-out/fan-in pattern
