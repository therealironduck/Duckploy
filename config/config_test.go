package config

import (
	"Duckploy/helper"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestThatItCanReadTheConfigFromAFilePath(t *testing.T) {
	t.Parallel()

	helper.AppFs = afero.NewMemMapFs()

	err := helper.AppFs.MkdirAll("some/path/to", os.ModePerm)
	require.NoError(t, err)

	err = afero.WriteFile(helper.AppFs, "some/path/to/config.json", []byte(helper.SimpleJsonConfig), os.ModePerm)
	require.NoError(t, err)

	result, err := ReadConfig("some/path/to/config.json")
	require.NoError(t, err)

	require.Len(t, result.Hosts, 1)
	require.Equal(t, "some-host", result.Hosts[0].Hostname)
	require.Equal(t, "ducky", result.Hosts[0].SSHUser)
	require.Equal(t, "secret123", result.Hosts[0].SSHPassword)
	require.Equal(t, "/some/path", result.Hosts[0].Path)

	require.Len(t, result.Steps, 2)
	require.Equal(t, "command", result.Steps[0].Action)
	require.Equal(t, "npm install", result.Steps[0].Command)

	require.Equal(t, "command", result.Steps[1].Action)
	require.Equal(t, "composer install", result.Steps[1].Command)
}

func TestErrorHandlingWhileReadingConfig(t *testing.T) {
	t.Parallel()

	t.Run("throws an error if the file does not exists", func(t *testing.T) {
		t.Parallel()

		helper.AppFs = afero.NewMemMapFs()

		_, err := ReadConfig("some/path/to/config.json")
		require.Error(t, err)
	})

	t.Run("throws an error if the file is not a valid json", func(t *testing.T) {
		t.Parallel()

		helper.AppFs = afero.NewMemMapFs()

		err := helper.AppFs.MkdirAll("some/path/to", os.ModePerm)
		require.NoError(t, err)

		err = afero.WriteFile(helper.AppFs, "some/path/to/config.json", []byte(`{`), os.ModePerm)
		require.NoError(t, err)

		_, err = ReadConfig("some/path/to/config.json")
		require.Error(t, err)
	})
}
