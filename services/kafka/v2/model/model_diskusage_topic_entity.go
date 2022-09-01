package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskusageTopicEntity struct {

	// 磁盘使用量。
	Size *string `json:"size,omitempty" xml:"size"`

	// topic名称。
	TopicName *string `json:"topic_name,omitempty" xml:"topic_name"`

	// 分区。
	TopicPartition *string `json:"topic_partition,omitempty" xml:"topic_partition"`

	// 磁盘使用量的占比。
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage"`
}

func (o DiskusageTopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskusageTopicEntity struct{}"
	}

	return strings.Join([]string{"DiskusageTopicEntity", string(data)}, " ")
}
