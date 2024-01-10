package storj

import (
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/storj/sub" // Import the sub package
)

var StorjCmd = &cobra.Command{
    Use:   "storj",
    Short: "Short description of storj",
    Long:  "Long description of storj",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    StorjCmd.AddCommand(sub.AddStorjCmd)
    StorjCmd.AddCommand(sub.GetStorjCmd)
    StorjCmd.AddCommand(sub.DeleteStorjCmd)
}
