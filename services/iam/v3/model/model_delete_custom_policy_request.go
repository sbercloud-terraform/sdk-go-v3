/*
 * IAM
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
type DeleteCustomPolicyRequest struct {
	RoleId string `json:"role_id"`
}

func (o DeleteCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomPolicyRequest", string(data)}, " ")
}
