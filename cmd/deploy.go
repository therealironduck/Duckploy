package cmd

import (
	"Duckploy/config"
	"Duckploy/helper"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [path]",
	Short: "Deploy the application!",
	Long:  `Load the Duckploy configuration and deploy it on the remote hosts.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		exists, _ := afero.Exists(helper.AppFs, args[0])
		if !exists {
			helper.Exitf(os.Stderr, "Duckploy configuration not found at: %s", args[0])
		}

		config, _ := config.ReadConfig(args[0])
		for _, host := range config.Hosts {
			fmt.Printf("Connecting to %s@%s:22 via password\n", host.SSHUser, host.Hostname)
			client, _ := helper.GetPasswordClient(host.SSHUser, host.SSHPassword, host.Hostname)

			for _, step := range config.Steps {
				fmt.Printf("-> Running `%s`\n", step.Command)
				client.Run(fmt.Sprintf("cd %s && %s", host.Path, step.Command))
				fmt.Printf("-> Done\n\n")
			}
		}
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
