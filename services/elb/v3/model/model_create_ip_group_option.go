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

// 创建IP地址组请求参数对象
type CreateIpGroupOption struct {
	// IP地址组的租户id
	ProjectId *string `json:"project_id,omitempty"`
	// IP地址组的描述信息
	Description *string `json:"description,omitempty"`
	// IP地址组的名称
	Name *string `json:"name,omitempty"`
	// IP地址组中包含的ip或网段列表。[]表示任意ip。
	IpList []CreateIpGroupIpOption `json:"ip_list"`
	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}
