package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复对象
type Restore struct {

	// 集群名称，要求唯一性，必须以字母开头并只包含字母、数字、中划线，下划线，长度为4~64个字符。
	Name string `json:"name" xml:"name"`

	// 指定子网ID，用于集群网络配置。默认值与原集群相同。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 指定安全组ID，用于集群网络配置。默认值与原集群相同。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 指定虚拟私有云ID，用于集群网络配置。默认值与原集群相同。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 指定集群可用区。默认值与原集群相同。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`

	// 指定集群服务端口
	Port *int32 `json:"port,omitempty" xml:"port"`

	PublicIp *PublicIp `json:"public_ip,omitempty" xml:"public_ip"`

	// 企业项目ID，对集群指定企业项目，如果未指定，则使用默认企业项目“default”的ID，即0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o Restore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Restore struct{}"
	}

	return strings.Join([]string{"Restore", string(data)}, " ")
}
