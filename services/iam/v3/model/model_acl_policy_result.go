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

//
type AclPolicyResult struct {
	// 允许访问的IP地址或网段。
	AllowAddressNetmasks []AllowAddressNetmasksResult `json:"allow_address_netmasks"`
	// 允许访问的IP地址区间。
	AllowIpRanges []AllowIpRangesResult `json:"allow_ip_ranges"`
}

func (o AclPolicyResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AclPolicyResult struct{}"
	}

	return strings.Join([]string{"AclPolicyResult", string(data)}, " ")
}
