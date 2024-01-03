package cmd

import (
    "os"
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/cmd/net"
    "github.com/dalecosta1/sinaloa-cli/cmd/info"
    "github.com/dalecosta1/sinaloa-cli/cmd/api"
)

var rootCmd = &cobra.Command{
    Use:   "sinaloa",
    Short: "The sinaloa cli",
    Long:  `Th sinaloa cli`,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func addSubcommandPalettes() {
    rootCmd.AddCommand(net.NetCmd)
    rootCmd.AddCommand(info.InfoCmd)
    rootCmd.AddCommand(api.ApiCmd)
}

func init() {
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    addSubcommandPalettes()
}
