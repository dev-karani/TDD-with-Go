package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people,", func(t *testing.T){
		got := Hello("edwin", "")
		want := "Hello, edwin"

		assertCorrectMessage(t, got , want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("","")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	//test for user passing spanish
	t.Run("in spanish", func(t *testing.T){
		got := Hello("elodie", "spanish")
		want := "Hola, elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got , want)
	}
}