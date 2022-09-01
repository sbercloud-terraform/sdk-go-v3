package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubsetsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 设备ID列表
	Items          *[]Device `json:"items,omitempty" xml:"items"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubsetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubsetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubsetsResponse", string(data)}, " ")
}
