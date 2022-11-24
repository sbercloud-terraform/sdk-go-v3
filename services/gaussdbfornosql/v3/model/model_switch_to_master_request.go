package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchToMasterRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o SwitchToMasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToMasterRequest struct{}"
	}

	return strings.Join([]string{"SwitchToMasterRequest", string(data)}, " ")
}
