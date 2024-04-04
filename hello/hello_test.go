package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike"

		if got != want {
			assertCorrectMessage(t, got, want)

		}
	})

	t.Run("saying Hello, World when there is an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("in spanish", func( t*testing.T){
		got := Hello("Mike", "Spanish")
		want := "Hola, Mike"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
