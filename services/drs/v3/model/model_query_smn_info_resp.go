package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 已录入SMN信息体
type QuerySmnInfoResp struct {

	// SMN信息
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty" xml:"subscriptions"`

	// 主题名称
	TopicName *string `json:"topic_name,omitempty" xml:"topic_name"`

	// 订阅延迟时间
	DelayTime *int64 `json:"delay_time,omitempty" xml:"delay_time"`

	// rto延迟时间
	RtoDelay *int64 `json:"rto_delay,omitempty" xml:"rto_delay"`

	// rpo延迟时间
	RpoDelay *int64 `json:"rpo_delay,omitempty" xml:"rpo_delay"`

	// 异常告警是否通知用户
	AlarmToUser *bool `json:"alarm_to_user,omitempty" xml:"alarm_to_user"`
}

func (o QuerySmnInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySmnInfoResp struct{}"
	}

	return strings.Join([]string{"QuerySmnInfoResp", string(data)}, " ")
}
