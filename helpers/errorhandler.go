package helpers

import (
    "encoding/json"
    "fmt"
    "github.com/dalecosta1/sinaloa-cli/models/messages/errors"
)

// HandleJSONError simplifies the creation and display of JSON-formatted error messages
func HandleResponseError(message, code string, err error) {
    errorMessage := fmt.Sprintf("%s: %s", message, err.Error())
    errorResponse := errors.NewErrorResponse(false, code, errorMessage)

    errorJsonResponse, jsonErr := json.MarshalIndent(errorResponse, "", "  ")
    if jsonErr != nil {
        fmt.Println("Error marshaling JSON:", jsonErr)
        return
    }

    fmt.Println(string(errorJsonResponse))
}