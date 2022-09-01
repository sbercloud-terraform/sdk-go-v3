package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRuleActionResponse struct {

	// 规则动作ID，用于唯一标识一条规则动作，在创建规则动作时由物联网平台分配获得，创建时无需携带，由平台统一分配唯一的action_id。
	ActionId *string `json:"action_id,omitempty" xml:"action_id"`

	// 规则动作对应的的规则触发条件ID。
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 资源空间ID。
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 规则动作的类型，取值范围： - HTTP_FORWARDING：HTTP服务消息类型。 - DIS_FORWARDING：转发DIS服务消息类型。 - OBS_FORWARDING：转发OBS服务消息类型。 - AMQP_FORWARDING：转发AMQP服务消息类型。 - DMS_KAFKA_FORWARDING：转发kafka消息类型。
	Channel *string `json:"channel,omitempty" xml:"channel"`

	ChannelDetail  *ChannelDetail `json:"channel_detail,omitempty" xml:"channel_detail"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateRuleActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleActionResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleActionResponse", string(data)}, " ")
}
