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
type UpdateHealthmonitorResponse struct {
	Healthmonitor *HealthmonitorV2Resp `json:"healthmonitor,omitempty"`
}

func (o UpdateHealthmonitorResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateHealthmonitorResponse", string(data)}, " ")
}
