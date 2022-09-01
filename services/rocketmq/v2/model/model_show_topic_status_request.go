package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTopicStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 主题名称。
	Topic string `json:"topic" xml:"topic"`
}

func (o ShowTopicStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRequest", string(data)}, " ")
}
