package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询重点指标。
type MetricDataPoints struct {

	// 统计方式。
	Statistics *[]StatisticValue `json:"statistics,omitempty" xml:"statistics"`

	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 时间序列单位。
	Unit *string `json:"unit,omitempty" xml:"unit"`
}

func (o MetricDataPoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataPoints struct{}"
	}

	return strings.Join([]string{"MetricDataPoints", string(data)}, " ")
}
