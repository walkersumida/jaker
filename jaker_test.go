package jaker

import (
	"testing"

	"github.com/walkersumida/prettyout"
)

func TestProfile(t *testing.T) {
	expected := 1

	t.Run("english fullname", func(t *testing.T) {
		got := len(Profile.EnFullName)
		if got < expected {
			t.Errorf(prettyout.Serror(expected, got))
		}
	})

	t.Run("email", func(t *testing.T) {
		got := len(Profile.Email)
		if got < expected {
			t.Errorf(prettyout.Serror(expected, got))
		}
	})

	t.Run("website", func(t *testing.T) {
		got := len(Profile.Website)
		if got < expected {
			t.Errorf(prettyout.Serror(expected, got))
		}
	})
}

func TestUuid(t *testing.T) {
	expected := 36
	got := len(Uuid)
	if got != expected {
			t.Errorf(prettyout.Serror(expected, got))
	}
}

func TestText(t *testing.T) {
	expected := "abcあいうabcあ"
	got := Text("abcあいう", 10)
	if got != expected {
			t.Errorf(prettyout.Serror(expected, got))
	}
}
