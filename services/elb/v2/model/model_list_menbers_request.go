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
type ListMenbersRequest struct {
	PoolId       string  `json:"pool_id"`
	Limit        *int32  `json:"limit,omitempty"`
	Marker       *string `json:"marker,omitempty"`
	PageReverse  *bool   `json:"page_reverse,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Address      *string `json:"address,omitempty"`
	ProtocolPort *int32  `json:"protocol_port,omitempty"`
	SubnetId     *string `json:"subnet_id,omitempty"`
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Weight       *int32  `json:"weight,omitempty"`
}

func (o ListMenbersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMenbersRequest", string(data)}, " ")
}
