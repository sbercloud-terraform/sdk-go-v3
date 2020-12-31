/*
 * ELB
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
type CreateL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateL7ruleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7ruleResponse struct{}"
	}

	return strings.Join([]string{"CreateL7ruleResponse", string(data)}, " ")
}
