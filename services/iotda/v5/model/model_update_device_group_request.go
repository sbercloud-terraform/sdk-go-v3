package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeviceGroupRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"Instance-Id,omitempty" xml:"Instance-Id"`

	// **参数说明**：设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。 **取值范围**：长度不超过36，十六进制字符串和连接符（-）的组合。
	GroupId string `json:"group_id" xml:"group_id"`

	Body *UpdateDeviceGroupDto `json:"body,omitempty" xml:"body"`
}

func (o UpdateDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupRequest", string(data)}, " ")
}
