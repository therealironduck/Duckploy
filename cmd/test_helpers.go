package cmd

import (
	"Duckploy/data"
	"bytes"
	"os"
	"testing"
)

func setupExitCodeTest(t *testing.T) (*int, func(), func()) {
	t.Helper()

	exitCode := 0
	exitFunc = func(code int) {
		exitCode = code
	}

	restoreFunc := func() {
		exitFunc = os.Exit
	}

	resetFunc := func() {
		exitCode = 0
	}

	return &exitCode, restoreFunc, resetFunc
}

func runCommandTest(t *testing.T, args ...string) {
	t.Helper()

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Command executing failed: %v", err)
	}
}

func pipeOutputTest(t *testing.T, run func()) (string, string, func()) {
	t.Helper()

	oldOut := os.Stdout
	oldErr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	resetFunc := func() {
		os.Stdout = oldOut
		os.Stderr = oldErr
	}

	run()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(rOut)
		outC <- buf.String()
	}()

	errC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(rErr)
		errC <- buf.String()
	}()

	err := wOut.Close()
	if err != nil {
		t.Fatalf("Cannot close output buffer: %v", err)
	}

	err = wErr.Close()
	if err != nil {
		t.Fatalf("Cannot close error buffer: %v", err)
	}

	out := <-outC
	outErr := <-errC

	return out, outErr, resetFunc
}

func mockConfigTest(t *testing.T, config data.Config) (*string, func(), func()) {
	t.Helper()

	var givenPath string

	old := readConfig
	restoreFunc := func() {
		readConfig = old
	}

	resetFunc := func() {
		givenPath = ""
	}

	readConfig = func(path string) data.Config {
		givenPath = path

		return config
	}

	return &givenPath, restoreFunc, resetFunc
}
