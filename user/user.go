// Package user provides a simple user struct with a Doer interface.
package user

import "github.com/zhang35/gomock-example/doer"

// User uses a Doer to do something.
type User struct {
	Doer doer.Doer
}

// Use uses the Doer to do some work.
func (u *User) Use(n int, s string) error {
	return u.Doer.DoSomething(n, s)
}
