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

// Request Object
type ShowLoadbalancersStatusRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancersStatusRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowLoadbalancersStatusRequest", string(data)}, " ")
}
