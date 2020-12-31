/*
 * CloudIDE
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckNameRequest struct {
	DisplayName string `json:"display_name"`
}

func (o CheckNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckNameRequest struct{}"
	}

	return strings.Join([]string{"CheckNameRequest", string(data)}, " ")
}
