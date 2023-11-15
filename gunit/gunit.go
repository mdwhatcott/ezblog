package gunit

import "testing"

func Run(f any, t *testing.T) {
	// TODO
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
