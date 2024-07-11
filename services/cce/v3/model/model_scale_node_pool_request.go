package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleNodePoolRequest Request Object
type ScaleNodePoolRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// 节点池ID
	NodepoolId string `json:"nodepool_id"`

	Body *ScaleNodePoolRequestBody `json:"body,omitempty"`
}

func (o ScaleNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleNodePoolRequest struct{}"
	}

	return strings.Join([]string{"ScaleNodePoolRequest", string(data)}, " ")
}
