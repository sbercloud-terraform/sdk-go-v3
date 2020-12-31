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

// 创建新集群信息
type CreateRequestBodyCceCreateInfo struct {
	// 集群节点数
	NodeNum *int32 `json:"node_num,omitempty"`
	// 集群节点规格
	NodeFlavor *string `json:"node_flavor,omitempty"`
	// CCE集群规格
	CceFlavor *string `json:"cce_flavor,omitempty"`
	// 节点初始密码
	InitNodePwd *string `json:"init_node_pwd,omitempty"`
	// 可用区
	Az *string `json:"az,omitempty"`
}

func (o CreateRequestBodyCceCreateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyCceCreateInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyCceCreateInfo", string(data)}, " ")
}
