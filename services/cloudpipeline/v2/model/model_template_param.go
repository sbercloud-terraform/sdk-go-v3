package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流水线参数
type TemplateParam struct {

	// 是否必须
	Required string `json:"required" xml:"required"`

	// 是否可见
	Visible string `json:"visible" xml:"visible"`

	// 流水线参数名字
	Name string `json:"name" xml:"name"`

	// 流水线参数值
	Value string `json:"value" xml:"value"`

	// 流水线参数描述
	Description string `json:"description" xml:"description"`

	// 流水线参数类型
	Paramtype string `json:"paramtype" xml:"paramtype"`

	// 流水线参数展示类型
	DisplayType string `json:"display_type" xml:"display_type"`

	// 流水线参数展示名字
	DisplayName string `json:"display_name" xml:"display_name"`

	// 是否静态参数
	IsStatic bool `json:"is_static" xml:"is_static"`

	// 是否默认参数
	IsDefault bool `json:"is_default" xml:"is_default"`

	// array类型数据
	Limits []ParamTypeLimits `json:"limits" xml:"limits"`

	// array类型数据
	Constraints []Constraint `json:"constraints" xml:"constraints"`
}

func (o TemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateParam struct{}"
	}

	return strings.Join([]string{"TemplateParam", string(data)}, " ")
}
