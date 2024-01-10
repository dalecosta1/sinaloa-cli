package env

import (
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/env/sub" // Import the sub package
)

// EnvCmd represents the env command
var EnvCmd = &cobra.Command{
    Use:   "env",
    Short: "This command is used to set and get env variables",
    Long:  "This command is used to set and get env variables",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    EnvCmd.AddCommand(sub.AddEnvCmd)
}
