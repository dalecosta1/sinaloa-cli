package sub  // Make sure this matches the directory structure

import (
    "fmt"
    "net/http"
    "time"
    "github.com/spf13/cobra"
)

var urlPath string
var client = &http.Client{
    Timeout: 2 * time.Second,
}

var PingCmd = &cobra.Command{
    Use:   "ping",
    Short: "A brief description of your command",
    Long:  ``,
    Run: func(cmd *cobra.Command, args []string) {
        if resp, err := ping(urlPath); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(resp)
        }
    },
}

func ping(domain string) (int, error) {
    url := "http://" + domain
    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
        return 0, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()  // Ensure the response body is closed

    return resp.StatusCode, nil
}

func init() {
    PingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")
    if err := PingCmd.MarkFlagRequired("url"); err != nil {
        fmt.Println(err)
    }
}
