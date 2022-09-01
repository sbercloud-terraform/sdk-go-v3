package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额资源
type PolicyInstanceResources struct {

	// 查询配额的类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 已使用的配额数量。
	Used *int32 `json:"used,omitempty" xml:"used"`

	// 配额总数量。
	Quota *int32 `json:"quota,omitempty" xml:"quota"`

	// 配额上限。
	Max *int32 `json:"max,omitempty" xml:"max"`

	// 配额下限。
	Min *int32 `json:"min,omitempty" xml:"min"`
}

func (o PolicyInstanceResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyInstanceResources struct{}"
	}

	return strings.Join([]string{"PolicyInstanceResources", string(data)}, " ")
}
