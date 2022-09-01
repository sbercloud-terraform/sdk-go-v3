package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 负载均衡器
type ModifyLb struct {
	LbaasListener *LbaasListener `json:"lbaas_listener,omitempty" xml:"lbaas_listener"`

	// 经典型负载均衡器信息
	Listener *string `json:"listener,omitempty" xml:"listener"`

	// 负载均衡器迁移失败原因。
	FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason"`

	// 负载均衡器迁移失败详情。
	FailedDetails *string `json:"failed_details,omitempty" xml:"failed_details"`
}

func (o ModifyLb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLb struct{}"
	}

	return strings.Join([]string{"ModifyLb", string(data)}, " ")
}
