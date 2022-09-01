package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务的动态属性
type SvcSpec struct {

	// 虚拟服务IP地址
	ClusterIp *string `json:"cluster_ip,omitempty" xml:"cluster_ip"`

	// 外部IP列表 --- 暂不支持
	ExternalIps *[]string `json:"external_ips,omitempty" xml:"external_ips"`

	// 外部域名 --- 暂不支持
	ExternalName *string `json:"external_name,omitempty" xml:"external_name"`

	// 服务需要暴露的端口列表
	Ports *[]SvcPort `json:"ports,omitempty" xml:"ports"`

	// 标签选择器，将选择具有指定Label标签的Pod作为管理范围
	Selector map[string]string `json:"selector" xml:"selector"`

	// 服务的类型
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o SvcSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SvcSpec struct{}"
	}

	return strings.Join([]string{"SvcSpec", string(data)}, " ")
}
