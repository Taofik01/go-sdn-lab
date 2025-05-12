package programs_test

import (
	"testing"

	"github.com/Taofik01/go-sdn-lab/programs"
)

func TestProgram(t *testing.T) {
	// This is a simple test function that checks if the Programs function returns the expected string.
	if programs.Programs() != "Go Gophers , let's code!" {
		t.Fatalf("Expected LFG, but got %s", programs.Programs())
	}
	expected := "Go Gophers , let's code!"
	actual := programs.Programs()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}

}
