/*
LESSON 14: Channels

Channels are typed conduits that let goroutines communicate safely. They are
central to Go's concurrency model and enable patterns like pipelines and worker pools.

Types of channels:
- Unbuffered: send and receive synchronize goroutines (blocking until the other side is ready).
- Buffered: make(chan T, n) allows up to n queued values without blocking the sender.
- Directional: chan<- T (send-only) and <-chan T (receive-only) are used in function signatures to clarify intent.

Common operations:
- ch := make(chan int)        // create
- ch <- v                     // send
- v := <-ch                   // receive
- close(ch)                   // close channel (sends panic if you try to send after close)

Patterns & tips:
- Use range ch to receive until the channel is closed.
- Use select to wait on multiple channel operations and implement timeouts.
- Avoid closing channels from the receiver side; the sender should close when done.

Practice tip: build a simple pipeline of stages: producer -> squarer -> printer.
*/

package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Basic channel creation and communication
	// TODO: Create an unbuffered channel and send/receive data
	fmt.Println("============================================================")
	fmt.Println("=== EXERCISE 1: Basic Channels ===")
	basicChannelExample()

	// Exercise 2: Buffered channels
	// TODO: Create buffered channels with different sizes
	// TODO: Understand blocking behavior
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 2: Buffered Channels ===")
	bufferedChannelExample()

	// Exercise 3: Channel directions
	// TODO: Create functions with send-only and receive-only channels
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 3: Channel Directions ===")
	channelDirectionsExample()

	// Exercise 4: Closing channels and range
	// TODO: Close channels and iterate using range
	// TODO: Check if channel is closed using the comma ok idiom
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 4: Closing Channels ===")
	closingChannelsExample()

	// Exercise 5: Select statement
	// TODO: Use select for non-blocking channel operations
	// TODO: Handle multiple channels simultaneously
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 5: Select Statement ===")
	selectExample()

	// Exercise 6: Timeout with select
	// TODO: Implement timeouts using select and time.After
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 6: Timeouts ===")
	timeoutExample()

	// Exercise 7: Pipeline pattern
	// TODO: Create a pipeline of goroutines connected by channels
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 7: Pipeline ===")
	pipelineExample()

	// Exercise 8: Fan-out/Fan-in pattern
	// TODO: Distribute work to multiple workers and collect results
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 8: Fan-out/Fan-in ===")
	fanOutFanInExample()

	// Exercise 9: Worker pool with channels
	// TODO: Implement a worker pool using channels for job distribution
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 9: Worker Pool ===")
	workerPoolChannelExample()

	// Exercise 10: Rate limiting with channels
	// TODO: Implement rate limiting using channels and time.Tick
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 10: Rate Limiting ===")
	rateLimitingExample()

	// Exercise 11: Channel synchronization patterns
	// TODO: Use channels for synchronization (done channels)
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 11: Synchronization ===")
	synchronizationExample()

	// Exercise 12: Nil channels and their behavior
	// TODO: Understand nil channel behavior in select statements
	fmt.Println("============================================================")
	fmt.Println("\n=== EXERCISE 12: Nil Channels ===")
	nilChannelExample()
}

func basicChannelExample() {
	// TODO: Create unbuffered channel
	// TODO: Send and receive data between goroutines
	fmt.Println("TODO: Implement basic channel communication")
}

func bufferedChannelExample() {
	// TODO: Create buffered channels
	// TODO: Show difference in blocking behavior
	fmt.Println("TODO: Implement buffered channel examples")
}

func channelDirectionsExample() {
	// TODO: Implement functions with directional channels
	fmt.Println("TODO: Implement channel direction examples")
}

func closingChannelsExample() {
	// TODO: Close channels and use range to iterate
	// TODO: Check for closed channels
	fmt.Println("TODO: Implement closing channels examples")
}

func selectExample() {
	// TODO: Use select statement with multiple channels
	// TODO: Include default case for non-blocking behavior
	fmt.Println("TODO: Implement select statement examples")
}

func timeoutExample() {
	// TODO: Implement timeout using select and time.After
	fmt.Println("TODO: Implement timeout examples")
}

func pipelineExample() {
	// TODO: Create a processing pipeline with multiple stages
	// Each stage should be a goroutine connected by channels
	fmt.Println("TODO: Implement pipeline pattern")
}

func fanOutFanInExample() {
	// TODO: Split work among multiple goroutines
	// TODO: Collect results from all workers
	fmt.Println("TODO: Implement fan-out/fan-in pattern")
}

func workerPoolChannelExample() {
	// TODO: Create a pool of worker goroutines
	// TODO: Distribute jobs through a channel
	fmt.Println("TODO: Implement worker pool with channels")
}

func rateLimitingExample() {
	// TODO: Limit rate of operations using channels
	// TODO: Use time.Tick or time.NewTicker
	fmt.Println("TODO: Implement rate limiting")
}

func synchronizationExample() {
	// TODO: Use done channels for synchronization
	// TODO: Signal completion between goroutines
	fmt.Println("TODO: Implement synchronization patterns")
}

func nilChannelExample() {
	// TODO: Show behavior of nil channels in select
	// TODO: Use nil channels to disable cases in select
	fmt.Println("TODO: Implement nil channel examples")
}

// Helper functions to implement:

// TODO: producer(ch chan<- int) - sends data to channel
// TODO: consumer(ch <-chan int) - receives data from channel
// TODO: squarer(in <-chan int, out chan<- int) - pipeline stage
// TODO: merge(cs ...<-chan int) <-chan int - fan-in function
// TODO: split(in <-chan int) (<-chan int, <-chan int) - fan-out function
