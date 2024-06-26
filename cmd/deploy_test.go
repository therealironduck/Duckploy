package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestThatItThrowsAnErrorIfPathWasNotFoundAndNonInteractive(t *testing.T) {
// 	t.Parallel()

// 	exitCode, resetFuncExitCode, _ := setupExitCodeTest(t)
// 	defer resetFuncExitCode()

// 	_, out, resetFuncOutput := pipeOutputTest(t, func() {
// 		runCommandTest(t, "deploy", "-n")
// 	})
// 	defer resetFuncOutput()

// 	assert.Equal(t, 1, *exitCode, "Exit code should be 1")
// 	assert.Contains(t, out, "Duckploy configuration not found at: duckploy.json")
// }

func TestThatThatItCanConnectViaSsh(t *testing.T) {
	t.Parallel()

	exitCode, restoreFuncExitCode, resetFuncExitCode := setupExitCodeTest(t)
	t.Cleanup(restoreFuncExitCode)

	t.Run("via password", func(t *testing.T) {
		t.Parallel()
		t.Cleanup(resetFuncExitCode)

		out, err, restoreFuncOutput := pipeOutputTest(t, func() {
			runCommandTest(t, "deploy")
		})
		t.Cleanup(restoreFuncOutput)

		assert.Empty(t, err)
		assert.Empty(t, out)

		assert.Equal(t, 0, *exitCode)
	})
}
