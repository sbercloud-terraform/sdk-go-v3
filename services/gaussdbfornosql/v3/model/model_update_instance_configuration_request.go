package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceConfigurationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateInstanceConfigurationRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequest", string(data)}, " ")
}
