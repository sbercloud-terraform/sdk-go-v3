package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConfigMapRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 配置项ID
	ConfigmapId string `json:"configmap_id" xml:"configmap_id"`
}

func (o DeleteConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigMapRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigMapRequest", string(data)}, " ")
}
