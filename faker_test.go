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

func TestUuid(t *testing.T) {
	expected := 36
	got := len(Uuid)
	if got != expected {
		t.Errorf("\nexpected: %d\n     got: %d", expected, got)
	}
}

func TestText(t *testing.T) {
	expected := "abcあいうabcあ"
	got := Text("abcあいう", 10)
	if got != expected {
		t.Errorf("\nexpected: %s\n     got: %s", expected, got)
	}
}
