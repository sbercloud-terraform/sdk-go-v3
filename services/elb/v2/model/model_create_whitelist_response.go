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
type CreateWhitelistResponse struct {
	Whitelist *WhitelistV2Resp `json:"whitelist,omitempty"`
}

func (o CreateWhitelistResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateWhitelistResponse", string(data)}, " ")
}
