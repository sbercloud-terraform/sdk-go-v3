package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNotificationResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 订阅管理列表
	Items          *[]NotificationResponseBody `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationResponse", string(data)}, " ")
}
