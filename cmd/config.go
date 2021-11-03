package cmd

import "github.com/spf13/cobra"

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "The 'config' subcommand is for use in management of configuration.",
	Long: `The 'config' subcommand is for use in management of configuration. It can be used, in combination with the
other subcommands 'add', 'update', 'view', and 'delete', . For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringP("key", "k", "", "The key for the key value set to add to the configuration.")
	configCmd.PersistentFlags().StringP("value", "v", "", "The value for the key value set to add to the configuration.")
}
