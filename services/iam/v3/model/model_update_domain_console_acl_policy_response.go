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
type UpdateDomainConsoleAclPolicyResponse struct {
	ConsoleAclPolicy *AclPolicyResult `json:"console_acl_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o UpdateDomainConsoleAclPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyResponse", string(data)}, " ")
}
