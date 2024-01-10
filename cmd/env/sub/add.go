package sub

import (
    "os"

    "github.com/spf13/cobra"
)

var (
    key string
    value string
)

var AddEnvCmd = &cobra.Command{
    Use:   "add",
    Short: "Add env variable",
    Long:  "Add env variable. Example: sinaloa env add -k=MY_VAR -v=MY_VALUE",
    Run: func(cmd *cobra.Command, args []string) {
        // Set env variable
        os.Setenv(key, value)
    },
}

func init() {
    AddEnvCmd.Flags().StringVarP(&key, "key", "k", "", "Env variable name to add")
    AddEnvCmd.Flags().StringVarP(&value, "value", "v", "", "Value of the variable name to add")
}
