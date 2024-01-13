package be 

import (
    "encoding/json"
	"fmt"
	"github.com/dalecosta1/sinaloa-cli/helpers"
	"github.com/dalecosta1/sinaloa-cli/storj"
)

// Implement fucntion to renew haproxy SSL certificates
func RenewHaproxySSLCertificatesStorj(storjPathCert string, storjPathKey string, secret string) []byte {
	// Implement renew haproxy SSL certificates
	return helpers.HandleResponse("Renew haproxy SSL certificates", "200", struct{}{})
}
