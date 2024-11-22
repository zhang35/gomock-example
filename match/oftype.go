package match

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// gomock/matchers.go excerpt
// type Matcher interface {
//     Matches(x interface{}) bool
//     String() string
// }

type ofType struct{ t string }

// OfType returns a gomock.Matcher that matches an interface{} value if it is of the specified type.
func OfType(t string) gomock.Matcher {
	return &ofType{t}
}

func (o *ofType) Matches(x interface{}) bool {
	return reflect.TypeOf(x).String() == o.t
}

func (o *ofType) String() string {
	return "is of type " + o.t
}
