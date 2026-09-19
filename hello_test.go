package main

import "testing"

func TestHelloWorld(t *testing.T) {
    if got, want := "hello", "hello"; got != want {
        t.Fatalf("got %q, want %q", got, want)
    }
}
