package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSqlFilterRuleRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点id。
	NodeId string `json:"node_id"`

	// SQL限流类型，取值为SELECT、UPDATE、DELETE，不区分大小写；若不传则默认查询所有类型的限流规则。
	Type *string `json:"type,omitempty"`
}

func (o ShowSqlFilterRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlFilterRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlFilterRuleRequest", string(data)}, " ")
}
