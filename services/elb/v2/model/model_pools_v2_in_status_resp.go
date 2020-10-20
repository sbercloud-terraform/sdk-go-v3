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

// 后端云服务器对象列表，用于状态树中
type PoolsV2InStatusResp struct {
	// 后端云服务器组ID
	Id string `json:"id"`
	// 后端云服务器组名称
	Name string `json:"name"`
	// 后端云服务器组关联的后端云服务器列表
	Members []MembersV2InStatusResp `json:"members"`
	// 后端云服务器组的操作状态；该字段为预留字段，暂未启用。默认为ONLINE。
	OperatingStatus string `json:"operating_status"`
	// 后端云服务器组的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。
	ProvisioningStatus string                        `json:"provisioning_status"`
	Healthmonitor      *HealthmonitorsV2InStatusResp `json:"healthmonitor"`
}

func (o PoolsV2InStatusResp) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PoolsV2InStatusResp", string(data)}, " ")
}
