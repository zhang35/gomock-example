// Package doer provides an interface for doing something.
package doer

// Doer interface has a method that does something.
//
//go:generate mockgen -source=doer.go -package mocks -destination mocks/mock_doer.go
type Doer interface {
	DoSomething(int, string) error
}
