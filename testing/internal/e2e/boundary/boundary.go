// Package boundary provides methods for commonly used boundary actions that are used in end-to-end tests.
package boundary

import (
	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/authmethods"
)

// AuthenticateCliOutput parses the json response from running `boundary authenticate`
type AuthenticateCliOutput struct {
	Item     *authmethods.AuthenticateResult
	response *api.Response
}

// AuthMethodInfo parses auth method info in the json response from running `boundary database init`
type AuthMethodInfo struct {
	AuthMethodId string `json:"auth_method_id"`
	LoginName    string `json:"login_name"`
	Password     string `json:"password"`
}

// DbInitInfo parses the json response from running `boundary database init`
type DbInitInfo struct {
	AuthMethod AuthMethodInfo `json:"auth_method"`
}

// CliError parses the Stderr from running a boundary command
type CliError struct {
	Status int `json:"status"`
}