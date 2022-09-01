package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerGroup struct {

	// 是否启用。
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 是否广播。
	Broadcast *bool `json:"broadcast,omitempty" xml:"broadcast"`

	// 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty" xml:"brokers"`

	// 消费组名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 最大重试次数。
	RetryMaxTime float32 `json:"retry_max_time,omitempty" xml:"retry_max_time"`

	// 是否重头消费。
	FromBeginning *bool `json:"from_beginning,omitempty" xml:"from_beginning"`
}

func (o ConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroup struct{}"
	}

	return strings.Join([]string{"ConsumerGroup", string(data)}, " ")
}
