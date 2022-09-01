package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEventResponse struct {

	// 测试事件ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 测试事件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 测试事件content。
	Content *string `json:"content,omitempty" xml:"content"`

	// 上次修改的时间。
	LastModified   float32 `json:"last_modified,omitempty" xml:"last_modified"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventResponse struct{}"
	}

	return strings.Join([]string{"ShowEventResponse", string(data)}, " ")
}
