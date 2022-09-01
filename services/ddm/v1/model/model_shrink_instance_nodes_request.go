package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShrinkInstanceNodesRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ReduceRequest `json:"body,omitempty" xml:"body"`
}

func (o ShrinkInstanceNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodesRequest", string(data)}, " ")
}
