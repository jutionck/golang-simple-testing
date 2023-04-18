package service

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting() // "Hello
	expected := "Hello"
	if got != expected {
		panic("Wrong output should be Hello")
	}
}

func TestHai(t *testing.T) {
	got := Greeting() // "Hello
	expected := "Hello"
	if got != expected {
		panic("Wrong output should be Hello")
	}
}
