package course

import "fmt"

// Exercise describes a single exercise entry in the course registry.
type Exercise struct {
	ID    string
	Title string
	Impl  bool
	Run   func() string // returns short result or status
}

// Registry holds all exercises.
var Registry = []Exercise{}

// Register adds an exercise to the global registry.
func Register(e Exercise) {
	Registry = append(Registry, e)
}

// Report runs implemented exercises and returns a formatted string report.
func Report() string {
	out := "Course Report:\n"
	for _, e := range Registry {
		status := "PENDING"
		if e.Impl {
			status = "IMPLEMENTED"
		}
		out += fmt.Sprintf("- [%s] %s: %s\n", status, e.ID, e.Title)
	}
	out += "\nResults of implemented exercises:\n"
	for _, e := range Registry {
		if e.Impl && e.Run != nil {
			out += fmt.Sprintf("%s -> %s\n", e.ID, e.Run())
		}
	}
	return out
}
