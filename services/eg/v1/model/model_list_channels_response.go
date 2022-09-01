package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListChannelsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本页数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 对象列表
	Items          *[]ChannelInfo `json:"items,omitempty" xml:"items"`
	HttpStatusCode int            `json:"-"`
}

func (o ListChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListChannelsResponse", string(data)}, " ")
}
