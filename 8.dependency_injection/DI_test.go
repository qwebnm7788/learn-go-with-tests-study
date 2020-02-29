package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jaewon")

	got := buffer.String()
	want := "Hello, Jaewon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
