package cmd

import (
	"github.com/spf13/cobra"
	"path/filepath"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [path]",
	Short: "Deploy the application!",
	Long:  `Load the Duckploy configuration and deploy it on the remote hosts.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, _ []string) {
		path, _ := filepath.Abs("example/duckploy.json")
		readConfig(path)

		// exitWithErrorf("Duckploy configuration not found at: duckploy.json")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().BoolP(
		"non-interactive",
		"n",
		false,
		"Disables all interactive prompts",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
