package faker

import "testing"

func TestFirstName(t *testing.T) {
	expected := "Walker"
	got := FirstName()
	if got != expected {
		t.Errorf("\nexpected: %s\n     got: %s", expected, got)
	}
}

func TestLastName(t *testing.T) {
	expected := "Sumida"
	got := LastName()
	if got != expected {
		t.Errorf("\nexpected: %s\n     got: %s", expected, got)
	}
}
