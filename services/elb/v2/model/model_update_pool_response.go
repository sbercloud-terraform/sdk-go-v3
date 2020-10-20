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
type UpdatePoolResponse struct {
	Pool *PoolV2Resp `json:"pool,omitempty"`
}

func (o UpdatePoolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePoolResponse", string(data)}, " ")
}
