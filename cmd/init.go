package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Initialize a new Duckploy project!",
	Long: `This command generates a new Duckploy template file, which can be used
as a starter point for deployments!`,
	Run: func(_ *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			exitWithErrorf("Failed to get absolut path: %v", err)
		}

		fileInfo, err := os.Stat(absPath)
		if err != nil {
			exitWithErrorf("Failed to get file info: %v", err)
		}

		if !fileInfo.IsDir() {
			exitWithErrorf("Path is not a directory: %s", absPath)
		}

		fmt.Println("Using path: " + absPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func exitWithErrorf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}
