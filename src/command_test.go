package main

import (
	"fmt"
	"testing"
)

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
	fmt.Println("Passed")
}

func TestSplitCommand(t *testing.T) {
	got := clearHistory()
	want := "History has been cleared ✔\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	fmt.Println("Passed")
}

func testCaptureOutput(t *testing.T) {
	got, err := captureOutput("cat ci/data")
	want := "test\n"
	if err != nil {
		t.Errorf("got %q want %q", got, want)
	}
	if string(got) != want {
		t.Errorf("got %q want %q", got, want)
	}
	fmt.Println("Passed")
}
