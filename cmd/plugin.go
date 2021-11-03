package cmd

import (
	"github.com/spf13/cobra"
	"plugin-hashcorp/newApp/plugin"
)

func init() {
	pluginManager := plugin.NewManager()
	runnablePlugins := pluginManager.GetRunnablePlugins()
	for _, runnablePlugin := range runnablePlugins {
		command := &cobra.Command{
			Use:   runnablePlugin.Name,
			Short: runnablePlugin.Short,
			Run: func(cmd *cobra.Command, args []string) {
				argsRun := make([]string, 0)
				for _, arg := range runnablePlugin.Args {
					valArgs, _ := cmd.Flags().GetString(arg)
					argsRun = append(argsRun, valArgs)
				}
				runnablePlugin.Run(argsRun)
			},
		}
		rootCmd.AddCommand(command)
		for _, arg := range runnablePlugin.Args {
			command.PersistentFlags().String(arg, "", "")
		}
	}
}