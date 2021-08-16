package faker

import "testing"

func TestFirstName(t *testing.T) {
	expected := 1
	got := len(firstName().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestLastName(t *testing.T) {
	expected := 1
	got := len(lastName().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestProfile(t *testing.T) {
	expected := 1
	got := len(Profile.EnFullName)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}
