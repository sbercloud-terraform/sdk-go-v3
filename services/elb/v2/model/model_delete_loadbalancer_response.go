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
type DeleteLoadbalancerResponse struct {
}

func (o DeleteLoadbalancerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteLoadbalancerResponse", string(data)}, " ")
}
