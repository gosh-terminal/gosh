package main

import "testing"

func TestGetArg(t *testing.T) {
	got := getArg("ls data")
	want := "data"
	got1 := getArg("cd ")
	want1 := ""
	if got != want {
        t.Errorf("got %q want %q", got, want)
	}
	if got1 != want1 {
        t.Errorf("got %q want %q", got1, want1)
    }
}