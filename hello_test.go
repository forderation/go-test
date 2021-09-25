package main

import (
	"testing"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// separeate to test domain
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func Add(x, y int) int {
	return x + y
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

/**
Test Function Below
**/

func TestHello(t *testing.T) {
	got := Hello("World", "")
	want := "Hello, World" // expected result
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCompose(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // find the correct line number that have wrong
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		// test pass
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// test fail
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("", spanish)
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}
