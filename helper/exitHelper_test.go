package helper

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItCanExitTheApplicationAndPrintAMessage(t *testing.T) {
	t.Parallel()

	var exitCode int
	ExitFunc = func(code int) {
		exitCode = code
	}
	defer func() {
		ExitFunc = os.Exit
	}()

	var b bytes.Buffer

	Exitf(&b, "This is an error message: %s", "and this is a parameter")

	require.Equal(t, "This is an error message: and this is a parameter\n", b.String())
	require.Equal(t, 1, exitCode)
}
