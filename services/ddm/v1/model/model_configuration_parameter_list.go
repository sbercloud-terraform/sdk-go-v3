package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationParameterList struct {

	// 参数名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 参数值。
	Value *string `json:"value,omitempty" xml:"value"`

	// 是否需要重启实例。
	NeedRestart *string `json:"need_restart,omitempty" xml:"need_restart"`

	// 参数是否只读。
	ReadOnly *string `json:"read_only,omitempty" xml:"read_only"`

	// 参数取值范围。
	ValueRange *string `json:"value_range,omitempty" xml:"value_range"`

	// 参数类型。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 参数描述。
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ConfigurationParameterList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterList struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterList", string(data)}, " ")
}
