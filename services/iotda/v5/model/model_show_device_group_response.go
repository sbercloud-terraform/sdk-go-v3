package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeviceGroupResponse struct {

	// 设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 设备组名称，单个资源空间下不可重复。
	Name *string `json:"name,omitempty" xml:"name"`

	// 设备组描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 父设备组ID，该设备组的父设备组ID。
	SuperGroupId   *string `json:"super_group_id,omitempty" xml:"super_group_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupResponse", string(data)}, " ")
}
