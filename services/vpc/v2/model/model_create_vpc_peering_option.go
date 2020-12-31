/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// peering对象
type CreateVpcPeeringOption struct {
	// 功能说明：对等连接名称 取值范围：支持1~64个字符
	Name           string   `json:"name"`
	RequestVpcInfo *VpcInfo `json:"request_vpc_info"`
	AcceptVpcInfo  *VpcInfo `json:"accept_vpc_info"`
}

func (o CreateVpcPeeringOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringOption struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringOption", string(data)}, " ")
}
