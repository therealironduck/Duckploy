package cmd

import (
	"Duckploy/data"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestThatItThrowsAnErrorIfPathWasNotFoundAndNonInteractive(t *testing.T) {
	t.Parallel()

	exitCode, resetFuncExitCode, _ := setupExitCodeTest(t)
	defer resetFuncExitCode()

	_, out, resetFuncOutput := pipeOutputTest(t, func() {
		runCommandTest(t, "deploy", "-n")
	})
	defer resetFuncOutput()

	assert.Equal(t, 1, *exitCode, "Exit code should be 1")
	assert.Contains(t, out, "Duckploy configuration not found at: duckploy.json")
}

func TestThatThatItCanConnectViaSsh(t *testing.T) {
	t.Parallel()

	exitCode, restoreFuncExitCode, resetFuncExitCode := setupExitCodeTest(t)
	defer restoreFuncExitCode()

	t.Run("via password", func(t *testing.T) {
		t.Parallel()

		defer resetFuncExitCode()

		path, restoreFuncConfig, _ := mockConfigTest(t, data.Config{
			Hosts: []data.Host{
				{Hostname: "fake-host", SshUser: "fakey", SshPassword: "my-fake-pw", Path: "/home/fakey/example"},
			},
		})
		defer restoreFuncConfig()

		_, _, restoreFuncOutput := pipeOutputTest(t, func() {
			runCommandTest(t, "deploy")
		})
		defer restoreFuncOutput()

		assert.Condition(t, func() (success bool) {
			return strings.HasSuffix(*path, "duckploy.json")
		})
		assert.Equal(t, 0, *exitCode)
	})

	// Basically: Load a ducktion.json file
	// Read the hosts section
	// Establish an ssh connection
	// Run all commands sequential
}
