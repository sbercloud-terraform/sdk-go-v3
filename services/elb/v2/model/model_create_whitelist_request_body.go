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

// This is a auto create Body Object
type CreateWhitelistRequestBody struct {
	Whitelist *CreateWhitelistV2Req `json:"whitelist"`
}

func (o CreateWhitelistRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateWhitelistRequestBody", string(data)}, " ")
}
