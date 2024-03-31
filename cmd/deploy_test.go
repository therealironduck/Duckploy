package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestDeployCmd(t *testing.T) {
	// Redirect output to buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute the command
	rootCmd.SetArgs([]string{"deploy"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Execute command with args failed: %v", err)
	}

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		buf.ReadFrom(r)
		outC <- buf.String()
	}()

	err := w.Close()
	if err != nil {
		t.Fatalf("Cannot close buffer: %v", err)
	}
	os.Stdout = old
	out := <-outC

	expected := "deploy called\n"
	if out != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, out)
	}
}
