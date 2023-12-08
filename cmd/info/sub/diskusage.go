package sub

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var (
	path string
)

var DiskUsageCmd = &cobra.Command{
	Use:   "disk-usage",
	Short: "Prints disk usage of a directory passed by arg.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the provided path is a directory
		fi, err := os.Stat(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if !fi.IsDir() {
			fmt.Println("Error: The specified path is not a directory.")
			return
		}

		// Get the size of the directory
		dirSize, err := getDirSize(path)
		if err != nil {
			fmt.Println("Error getting directory size:", err)
			return
		}

		// Calculate disk usage in gigabytes
		usage := du.NewDiskUsage(path)
		totalGB := float64(usage.Size()) / (1024 * 1024 * 1024)
		freeGB := float64(usage.Free()) / (1024 * 1024 * 1024)
		usedGB := float64(usage.Used()) / (1024 * 1024 * 1024)

		// Create a map for JSON representation
		diskUsageJSON := map[string]string{
			"TotalGb":        fmt.Sprintf("%.2f", totalGB),
			"FreeGb":         fmt.Sprintf("%.2f", freeGB),
			"UsedGb":         fmt.Sprintf("%.2f", usedGB),
			"DirSizeMb":      fmt.Sprintf("%.2f", float64(dirSize) / (1024 * 1024)),
			"DirSizeRawByte": fmt.Sprintf("%d", dirSize),
		}		

		// Convert map to JSON
		jsonData, err := json.Marshal(diskUsageJSON)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// Print JSON data
		fmt.Println(string(jsonData))
	},
}

func init() {
	DiskUsageCmd.Flags().StringVarP(&path, "path", "p", "", "Path of the folder to check disk usage.")

	if err := DiskUsageCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}
}

func getDirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size, err
}
