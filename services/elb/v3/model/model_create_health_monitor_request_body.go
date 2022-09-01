package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateHealthMonitorRequestBody struct {
	Healthmonitor *CreateHealthMonitorOption `json:"healthmonitor" xml:"healthmonitor"`
}

func (o CreateHealthMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequestBody", string(data)}, " ")
}
