/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 节点组织
type NodeOrgs struct {
	// 组织名称
	Name *string `json:"name,omitempty"`
	// 组织目标节点数
	NodeCount *int32 `json:"node_count,omitempty"`
	// pvc名称
	PvcName *string `json:"pvc_name,omitempty"`
}

func (o NodeOrgs) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeOrgs struct{}"
	}

	return strings.Join([]string{"NodeOrgs", string(data)}, " ")
}
