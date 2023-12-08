package info

import (
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/info/sub" // Import the sub package
)

var InfoCmd = &cobra.Command{
    Use:   "info",
    Short: "A brief description of your command",
    Long:  ``,
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    InfoCmd.AddCommand(sub.DiskUsageCmd) // Use the DiskUsageCmd from the sub package
}
