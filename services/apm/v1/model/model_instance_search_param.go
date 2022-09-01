package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取实例信息列表入参
type InstanceSearchParam struct {

	// 环境id
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 当前页码
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 每页数据容量
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 关键字
	Keyword *string `json:"keyword,omitempty" xml:"keyword"`

	// 实例状态
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 是否返回计数结果
	ReturnCount *bool `json:"return_count,omitempty" xml:"return_count"`
}

func (o InstanceSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSearchParam struct{}"
	}

	return strings.Join([]string{"InstanceSearchParam", string(data)}, " ")
}
