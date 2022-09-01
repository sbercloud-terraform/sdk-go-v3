package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeviceTemplateRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 设备模板ID
	DeviceTemplateId string `json:"device_template_id" xml:"device_template_id"`
}

func (o DeleteDeviceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceTemplateRequest", string(data)}, " ")
}
