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

// Response Object
type UpdateDomainLoginPolicyResponse struct {
	LoginPolicy    *LoginPolicyResult `json:"login_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateDomainLoginPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyResponse", string(data)}, " ")
}
