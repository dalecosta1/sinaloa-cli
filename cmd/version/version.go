package version

import (
    "fmt"
    "encoding/json"
    "github.com/spf13/cobra"
    "github.com/dalecosta1/sinaloa-cli/models/messages/response"
    "github.com/dalecosta1/sinaloa-cli/models/messages/errors"
    "github.com/dalecosta1/sinaloa-cli/helpers"
)

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
    Use:   "version",
    Short: "Get the version of sinaloa-cli",
    Long:  "Get the version of sinaloa-cli. Example: sinaloa version",
    Run: func(cmd *cobra.Command, args []string) {
        // Load the config
	    helpers.LoadConfig()
        
        // Check if the version is set
        if helpers.AppConfig.VERSION == "" {
            // Create a 404 error response
            errorResponse := errors.NewErrorResponse(false, "404", "Version not found in environment")

            // Marshal the error response into JSON with indentation
            errorJsonResponse, err := json.MarshalIndent(errorResponse, "", "  ")
            if err != nil {
                fmt.Println("Error marshaling JSON:", err)
                return
            }

            // Print the indented JSON error response
            fmt.Println(string(errorJsonResponse))
            return
        }

        // Create the response
        versionResponse := response.NewResponse(true, "200", "", struct {
            Version string `json:"version"`
        }{
            Version: helpers.AppConfig.VERSION,
        })

        // Marshal the response into JSON with indentation
        jsonResponse, err := json.MarshalIndent(versionResponse, "", "  ")
        if err != nil {
            errorMessage := fmt.Sprintf("An error occurred: %s", err.Error())
            errorResponse := errors.NewErrorResponse(false, "500", errorMessage)
        
            errorJsonResponse, jsonErr := json.MarshalIndent(errorResponse, "", "  ")
            if jsonErr != nil {
                fmt.Println("Error marshaling the error message:", jsonErr)
                return
            }
        
            fmt.Println(string(errorJsonResponse))
            return
        }
        
        fmt.Println(string(jsonResponse))
    },
}

func init() {
    // Add your code here.
}
