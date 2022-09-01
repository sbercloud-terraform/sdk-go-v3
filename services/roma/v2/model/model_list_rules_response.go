package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRulesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 规则列表
	Items          *[]Rule `json:"items,omitempty" xml:"items"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRulesResponse", string(data)}, " ")
}
