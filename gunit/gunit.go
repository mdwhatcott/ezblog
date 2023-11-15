package gunit

import (
	"reflect"
	"strings"
	"testing"
)

func Run(f any, t *testing.T) {
	type_ := reflect.TypeOf(f)
	for x := 0; x < type_.NumMethod(); x++ {
		method := type_.Method(x).Name
		if strings.HasPrefix(method, "Test") {
			t.Run(method, func(t *testing.T) {
				fixture := reflect.New(type_.Elem())
				fixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(NewFixture(t)))
				fixture.MethodByName("Setup").Call(nil)
				fixture.MethodByName(method).Call(nil)
			})
		}
	}
}

type Fixture struct{ *testing.T }

func NewFixture(t *testing.T) *Fixture {
	return &Fixture{T: t}
}

type assertion func(actual any, expected ...any) string

func (this *Fixture) So(actual any, assert assertion, expected ...any) {
	this.Helper()
	if failure := assert(actual, expected...); failure != "" {
		this.Error(failure)
	}
}
