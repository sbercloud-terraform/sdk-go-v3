package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRoutingRuleRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"Instance-Id,omitempty" xml:"Instance-Id"`

	Body *AddRuleReq `json:"body,omitempty" xml:"body"`
}

func (o CreateRoutingRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutingRuleRequest", string(data)}, " ")
}
