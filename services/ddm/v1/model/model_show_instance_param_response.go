package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceParamResponse struct {

	// DDM参数最后更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// DDM实例参数信息列表的集合。
	ConfigurationParameter *[]ConfigurationParameterList `json:"configuration_parameter,omitempty" xml:"configuration_parameter"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 集合总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamResponse", string(data)}, " ")
}
