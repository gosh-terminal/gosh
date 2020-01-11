package test

import (
	"fmt"
	shell "gosh/pkg"
	"testing"
)

func TestGetArg(t *testing.T) {
	gotArray := []string{shell.GetArg("ls data"), shell.GetArg("cd ")}
	wantArray := []string{"data", ""}
	if gotArray[0] != wantArray[0] {
		t.Errorf("got %q want %q", gotArray[0], wantArray[0])
	}
	if gotArray[1] != wantArray[1] {
		t.Errorf("got %q want %q", gotArray[1], wantArray[1])
	}
	fmt.Println("Passed")
}

func TestSplitCommand(t *testing.T) {
	got := shell.ClearHistory()
	want := "History has been cleared ✔\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	fmt.Println("Passed")
}

func testCaptureOutput(t *testing.T) {
	got, err := shell.CaptureOutput("cat ci/data")
	want := "test\n"
	if err != nil {
		t.Errorf("got %q want %q", got, want)
	}
	if string(got) != want {
		t.Errorf("got %q want %q", got, want)
	}
	fmt.Println("Passed")
}
