package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 定时配置
type CronConfig struct {

	// 定时配置名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 定时表达式
	Cron *string `json:"cron,omitempty" xml:"cron"`

	// 拉起预留实例个数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 开始时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 失效时间戳
	ExpiredTime *int64 `json:"expired_time,omitempty" xml:"expired_time"`
}

func (o CronConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CronConfig struct{}"
	}

	return strings.Join([]string{"CronConfig", string(data)}, " ")
}
