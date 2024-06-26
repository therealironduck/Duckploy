package cmd

import (
	"Duckploy/helper"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorHandling(t *testing.T) {
	t.Parallel()

	exitCode, restoreFuncExitCode, resetFuncExitCode := setupExitCodeTest(t)
	t.Cleanup(restoreFuncExitCode)

	helper.AppFs = afero.NewMemMapFs()

	t.Run("test that it throws an error if the config file does not exist (non interactive)", func(t *testing.T) {
		t.Cleanup(resetFuncExitCode)

		_, outErr, resetFuncOutput := pipeOutputTest(t, func() {
			runCommandTest(t, "deploy", "non-existing-file.json", "-n")
		})
		t.Cleanup(resetFuncOutput)

		assert.Equal(t, 1, *exitCode, "Exit code should be 1")
		assert.Contains(t, outErr, "Duckploy configuration not found at: non-existing-file.json")
	})
}

func TestThatItCanConnectViaSshAndExecuteCommands(t *testing.T) {
	exitCode, restoreFuncExitCode, resetFuncExitCode := setupExitCodeTest(t)
	t.Cleanup(restoreFuncExitCode)

	t.Run("via password", func(t *testing.T) {
		t.Cleanup(resetFuncExitCode)

		helper.AppFs = afero.NewMemMapFs()
		err := afero.WriteFile(helper.AppFs, "test-files/duckploy.json", []byte(helper.SimpleJsonConfig), 0777)
		require.NoError(t, err)

		out, outErr, resetFuncOutput := pipeOutputTest(t, func() {
			runCommandTest(t, "deploy", "test-files/duckploy.json")
		})
		t.Cleanup(resetFuncOutput)

		assert.Equal(t, 0, *exitCode, "Exit code should be 0")
		assert.Empty(t, outErr)
		assert.Contains(t, out, "Connecting to ducky@some-host:22 via password")
		// TODO
	})
}
