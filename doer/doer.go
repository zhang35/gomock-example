// Package doer provides an interface for doing something.
package doer

// Doer interface has a method that does something.
type Doer interface {
	DoSomething(int, string) error
}
