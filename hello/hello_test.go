package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Frost")
    want := "Hello, Frost"

    if got != want {
        t.Errorf("got '%q' want '%q'", got, want)
    }
}