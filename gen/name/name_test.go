package name_test

import (
	"testing"

	"github.com/walkersumida/jaker/gen/name"
)

func TestFirstName(t *testing.T) {
	expected := 1
	got := len(name.FirstGen().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestLastName(t *testing.T) {
	expected := 1
	got := len(name.LastGen().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}
