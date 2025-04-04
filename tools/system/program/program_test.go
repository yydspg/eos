package system

import (
	"os"
	"testing"
)

func TestGetProcessName(t *testing.T ) {
	originalArgs := os.Args
	defer func(){
		os.Args = originalArgs
	}()

	expected := "testprog"
	os.Args = []string{"/path/to/"+expected}
	now := getProcessName()
	if now != expected {
		t.Errorf("expected %s, but got %s", expected, now)
	}
}