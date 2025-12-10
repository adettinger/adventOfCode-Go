package testutils

import "testing"

func AssertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("found wrong error: got %q want %q", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got error %q but didnt expect one", got)
	}
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertInts(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertBool(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %t, Want %t", got, want)
	}
}
