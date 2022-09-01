package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警记录额外字段，仅针对事件监控告警场景所产生的告警记录信息。
type AdditionalInfo struct {

	// 该条告警记录对应的资源ID；如：22d98f6c-16d2-4c2d-b424-50e79d82838f。
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 该条告警记录对应的资源名称；如：ECS-Test01。
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 该条告警记录对应的事件监控ID，资源所产生的事件；如：ev16031292300990kKN8p17J。
	EventId *string `json:"event_id,omitempty" xml:"event_id"`
}

func (o AdditionalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalInfo struct{}"
	}

	return strings.Join([]string{"AdditionalInfo", string(data)}, " ")
}
