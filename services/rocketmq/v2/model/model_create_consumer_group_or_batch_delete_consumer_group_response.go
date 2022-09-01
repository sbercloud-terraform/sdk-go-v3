package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateConsumerGroupOrBatchDeleteConsumerGroupResponse struct {

	// 删除消费组的任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 创建成功的消费组名称。
	Name           *string `json:"name,omitempty" xml:"name"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConsumerGroupOrBatchDeleteConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupOrBatchDeleteConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupOrBatchDeleteConsumerGroupResponse", string(data)}, " ")
}
