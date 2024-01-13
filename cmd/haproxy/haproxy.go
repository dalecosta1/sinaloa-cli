package haproxy

import (
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/haproxy/sub"
)

var HaproxyCmd = &cobra.Command{
    Use:   "haproxy",
    Short: "Short description of haproxy",
    Long:  "Long description of haproxy",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    HaproxyCmd.AddCommand(sub.ReceiveHaproxyCmd)
}
