package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMembersDetailResponse struct {

	// 添加备份共享成员响应信息
	Members *[]Member `json:"members,omitempty" xml:"members"`

	// 备份共享成员数量
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMembersDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMembersDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailResponse", string(data)}, " ")
}
