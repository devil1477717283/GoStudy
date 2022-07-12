package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage:=func(t *testing.T,got,want string){
		t.Helper()
		if got!=want{
			t.Errorf("got '%q' want '%q'",got,want)
		} 
	}
	t.Run("saying hello to people",func(t *testing.T) {
		got :=Hello("people","english")
		want :="Hello,people"
		assertCorrectMessage(t,got,want)
	})
	t.Run("say hello world when an french string is supplied",func(t *testing.T){
		got :=Hello("","French")
		want :="Bonjour,World"
		assertCorrectMessage(t,got,want)
	})
	t.Run("say hello world when an empty string is supplied",func(t *testing.T){
		got :=Hello("","")
		want :="Hello,World"
		assertCorrectMessage(t,got,want)
	})
}