package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	port string
)

var StartApiCmd = &cobra.Command{
	Use:   "start",
	Short: "Start sinaloa api server.",
	Long:  `Start sinaloa api server.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the provided path is a directory
		fmt.Println("Starting api server...")
	},
}

func init() {
	StartApiCmd.Flags().StringVarP(&port, "port", "p", "", "Port of api server.")

	if err := StartApiCmd.MarkFlagRequired("port"); err != nil {
		fmt.Println(err)
	}
}

// func startApiServer(port string) (int64, error) {
// 	var size int64

// 	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() {
// 			size += info.Size()
// 		}
// 		return nil
// 	})

// 	return size, err
// }
