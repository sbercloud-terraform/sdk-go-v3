package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadSlowlogRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *DownloadSlowlogRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DownloadSlowlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogRequest struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogRequest", string(data)}, " ")
}
