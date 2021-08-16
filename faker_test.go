package faker

import "testing"

func TestFirstName(t *testing.T) {
	expected := 1
	got := len(PickUpFirstName().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestLastName(t *testing.T) {
	expected := 1
	got := len(PickUpLastName().En)
	if got < expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestProfile(t *testing.T) {
	expected := 1

	t.Run("english fullname", func(t *testing.T){
		got := len(Profile.EnFullName)
		if got < expected {
			t.Errorf("\nexpected: %d\n     got: %d", expected, got)
		}
	})

	t.Run("email", func(t *testing.T){
		got := len(Profile.Email)
		if got < expected {
			t.Errorf("\nexpected: %d\n     got: %d", expected, got)
		}
	})
}
