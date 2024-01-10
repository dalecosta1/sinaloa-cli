package net

import (
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/net/sub"
)

var NetCmd = &cobra.Command{
    Use:   "net",
    Short: "Net is a palette that contains network-based commands",
    Long:  ``,
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func init() {
    NetCmd.AddCommand(sub.PingCmd)  // Add the PingCmd from the sub package
}
