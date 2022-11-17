package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行部署任务时传递的动态参数
type KeyValueDo struct {

	// 执行部署任务时传递的参数名称
	Key *string `json:"key,omitempty"`

	// 执行部署任务时传递的参数值
	Value *string `json:"value,omitempty"`
}

func (o KeyValueDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyValueDo struct{}"
	}

	return strings.Join([]string{"KeyValueDo", string(data)}, " ")
}
