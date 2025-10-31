/*
LESSON 8: Maps

Maps store key-value pairs and are ideal for fast lookups. They’re reference types
and their iteration order is intentionally random.

Declaration ways:
- m := make(map[string]int)
- m := map[string]int{"a": 1, "b": 2}

Important notes:
- Accessing a missing key returns the zero value for the value type. Use the
	"comma ok" idiom to check presence: v, ok := m[key]
- Use delete(m, key) to remove an entry. Deleting a non-existent key is safe.
- Maps are not safe for concurrent writes; use sync.RWMutex or sync.Map for concurrency.

Practice tip: build a word frequency counter from a slice of strings using a map.
*/

package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	// Exercise 1: Map creation and basic operations
	// TODO: Create a map with string keys and int values using make()
	// m := map[string]int{"a": 1, "b": 2}
	// TODO: Add some key-value pairs
	// m["c"] = 3
	// TODO: Print the map and its length
	// fmt.Printf("m: %v\n", m)

	// Exercise 2: Map access and modification
	// TODO: Modify existing values
	// TODO: Try to access a key that doesn't exist
	// fmt.Printf("m['d']: %v\n", m["d"])

	// Exercise 4: The comma ok idiom
	// TODO: Use the two-value assignment to check if a key exists
	// v, ok := m["d"]
	// fmt.Printf("m['d'] exists: %b", ok)
	// fmt.Println(v)
	// TODO: Implement a function that safely gets a value with a default
	// fmt.Println(DefaultGet(m, "v", -1))
	// Exercise 5: Map iteration
	// TODO: Iterate over a map using range
	m := map[string]int{
		"Aleandro": 29,
		"Mario":    17,
	}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	fmt.Println("============================================================")

	// Exercise 6: Map deletion
	// TODO: Delete keys from a map using delete()
	// TODO: Try to delete a non-existent key (observe no error)
	delete(m, "Mario")
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	fmt.Println("============================================================")
	delete(m, "Franco")

	// Exercise 7: Nested maps
	// TODO: Create a map of maps (e.g., map[string]map[string]int)
	m2 := map[string]map[string]int{
		"Aleandro": {
			"Age":    29,
			"Status": 1,
		},
		"Filippo": {
			"Age":    19,
			"Status": 1,
		},
	}
	for key1, value1 := range m2 {
		fmt.Printf(key1)
		for key2, value2 := range value1 {
			fmt.Println(key2)
			fmt.Println(value2)
		}
	}
	fmt.Println("============================================================")

	// Exercise 8: Map as function parameters
	// TODO: Write functions that take maps as parameters
	fmt.Println("Map before function: ")
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	fmt.Println("Map after function: ")
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// TODO: Understand that maps are passed by reference
	m3 := mapFunction2(m)
	fmt.Println("Map after function that creates a new map: ")
	for key, value := range m3 {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	fmt.Println("============================================================")

	// Exercise 9: Practical exercises
	// TODO: Create a word counter that counts occurrences of words in a slice

	// TODO: Create a phone book (name -> phone number)
	// TODO: Implement a simple cache with get/set operations

	// Exercise 10: Advanced map operations
	// TODO: Create a function that merges two maps
	// TODO: Create a function that inverts a map (values become keys)
	// TODO: Create a function that filters map entries based on a condition
}

func DefaultGet(m map[string]int, index string, defaultValue int) int {
	if m[index] == 0 {
		return defaultValue
	} else {
		return m[index]
	}
}

func thatTakesAMap(asArgument map[string]int) {
	for _, value := range asArgument {
		value++
	}
}

func mapFunction2(m map[string]int) map[string]int {
	result := map[string]int{}
	for key, value := range m {
		result[key] = value + 1
	}
	return result
}

func workCounter(a []int) {
	m := map[string]int
	for 
}