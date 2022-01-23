package name_test

import (
	"testing"

	"github.com/walkersumida/jaker/gen/name"
	"github.com/walkersumida/prettyout"
)

func TestFirstName(t *testing.T) {
	expected := 1
	got := len(name.FirstGen().En)
	if got < expected {
		t.Errorf(prettyout.Serror(expected, got))
	}
}

func TestLastName(t *testing.T) {
	expected := 1
	got := len(name.LastGen().En)
	if got < expected {
		t.Errorf(prettyout.Serror(expected, got))
	}
}
