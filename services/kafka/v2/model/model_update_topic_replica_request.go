package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTopicReplicaRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// Topic名称。
	Topic string `json:"topic" xml:"topic"`

	Body *ResetReplicaReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateTopicReplicaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicReplicaRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicReplicaRequest", string(data)}, " ")
}
