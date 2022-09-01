package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricRequest2 struct {

	// 指标类型
	MetricType *string `json:"metric_type,omitempty" xml:"metric_type"`

	// 迭代ID
	SprintId *string `json:"sprint_id,omitempty" xml:"sprint_id"`

	Dividend *MetricRequest2Dividend `json:"dividend,omitempty" xml:"dividend"`

	// 指标分母过滤条件
	Divisor *interface{} `json:"divisor,omitempty" xml:"divisor"`
}

func (o MetricRequest2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricRequest2 struct{}"
	}

	return strings.Join([]string{"MetricRequest2", string(data)}, " ")
}
