/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVaultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVaultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVaultResponse struct{}"
	}

	return strings.Join([]string{"DeleteVaultResponse", string(data)}, " ")
}
