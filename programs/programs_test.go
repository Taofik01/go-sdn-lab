package programs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Taofik01/go-sdn-lab/programs"
)

func TestPrograms(t *testing.T) {
    // Simulate user input
    input := "4\n" // Simulate entering the number 4
    expectedOutput := "Enter a number: The number 4 is even.\n"

    // Redirect stdin
    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }()
    r, w, _ := os.Pipe()
    defer r.Close()
    w.WriteString(input)
    w.Close()
    os.Stdin = r

    // Redirect stdout
    var output bytes.Buffer
    oldStdout := os.Stdout
    defer func() { os.Stdout = oldStdout }()
    writer := io.MultiWriter(oldStdout, &output)
    os.Stdout = writer.(*os.File)

    // Call the function
    programs.Programs()

    // Check the output
    if output.String() != expectedOutput {
        t.Errorf("Expected output %q, but got %q", expectedOutput, output.String())
    }
}